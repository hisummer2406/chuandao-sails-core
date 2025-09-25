package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Snowflake ID 结构：64位
// +----------+------+------+------------+
// |  1bit    | 41bit| 10bit|   12bit    |
// +----------+------+------+------------+
// | 符号位   | 时间戳| 机器ID| 序列号     |
// +----------+------+------+------------+

const (
	// 时间戳位数
	timestampBits = 41
	// 机器ID位数
	workerIdBits = 10
	// 序列号位数
	sequenceBits = 12

	// 最大值计算
	maxWorkerId = (1 << workerIdBits) - 1 // 1023
	maxSequence = (1 << sequenceBits) - 1 // 4095

	// 位移量
	workerIdShift  = sequenceBits                // 12
	timestampShift = sequenceBits + workerIdBits // 22

	// 自定义时间起点 2024-01-01 00:00:00 UTC
	customEpoch = 1704067200000 // 毫秒时间戳
)

// OrderIdGenerator 订单ID生成器
type OrderIdGenerator struct {
	mutex         sync.Mutex
	workerId      int64
	lastTimestamp int64
	sequence      int64
}

// NewOrderIdGenerator 创建订单ID生成器
// workerId: 机器ID，范围 0-1023
func NewOrderIdGenerator(workerId int64) (*OrderIdGenerator, error) {
	if workerId < 0 || workerId > maxWorkerId {
		return nil, errors.New(fmt.Sprintf("worker id must be between 0 and %d", maxWorkerId))
	}

	return &OrderIdGenerator{
		workerId:      workerId,
		lastTimestamp: 0,
		sequence:      0,
	}, nil
}

// GenerateId 生成订单ID
func (g *OrderIdGenerator) GenerateId() (int64, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// 获取当前时间戳（毫秒）
	currentTimestamp := getCurrentTimestamp()

	// 时钟回拨检测
	if currentTimestamp < g.lastTimestamp {
		return 0, errors.New(fmt.Sprintf("clock moved backwards, refusing to generate id for %d milliseconds",
			g.lastTimestamp-currentTimestamp))
	}

	// 同一毫秒内序列号递增
	if currentTimestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & maxSequence

		// 序列号溢出，等待下一毫秒
		if g.sequence == 0 {
			currentTimestamp = g.waitNextMillis(currentTimestamp)
		}
	} else {
		// 新的毫秒，序列号重置
		g.sequence = 0
	}

	g.lastTimestamp = currentTimestamp

	// 组装64位ID
	// 时间戳部分：当前时间戳减去自定义起点
	timestampPart := (currentTimestamp - customEpoch) << timestampShift
	// 机器ID部分
	workerIdPart := g.workerId << workerIdShift
	// 序列号部分
	sequencePart := g.sequence

	id := timestampPart | workerIdPart | sequencePart

	return id, nil
}

// GenerateOrderId 生成带前缀的订单号字符串
func (g *OrderIdGenerator) GenerateOrderId() (string, error) {
	id, err := g.GenerateId()
	if err != nil {
		return "", err
	}

	// 格式：ORD + 19位数字ID
	return fmt.Sprintf("ORD%019d", id), nil
}

// GenerateOrderIdWithPrefix 生成自定义前缀的订单号
func (g *OrderIdGenerator) GenerateOrderIdWithPrefix(prefix string) (string, error) {
	id, err := g.GenerateId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%019d", prefix, id), nil
}

// BatchGenerate 批量生成订单ID
func (g *OrderIdGenerator) BatchGenerate(count int) ([]string, error) {
	if count <= 0 {
		return nil, errors.New("count must be greater than 0")
	}

	orderIds := make([]string, 0, count)

	for i := 0; i < count; i++ {
		orderId, err := g.GenerateOrderId()
		if err != nil {
			return nil, err
		}
		orderIds = append(orderIds, orderId)
	}

	return orderIds, nil
}

// ParseOrderId 解析订单ID，返回时间戳、机器ID、序列号
func (g *OrderIdGenerator) ParseOrderId(orderIdStr string) (timestamp int64, workerId int64, sequence int64, err error) {
	// 移除前缀，提取数字部分
	var id int64
	if len(orderIdStr) >= 3 && orderIdStr[:3] == "ORD" {
		_, err = fmt.Sscanf(orderIdStr[3:], "%d", &id)
	} else {
		_, err = fmt.Sscanf(orderIdStr, "%d", &id)
	}

	if err != nil {
		return 0, 0, 0, errors.New("invalid order id format")
	}

	// 解析各部分
	sequence = id & maxSequence
	workerId = (id >> workerIdShift) & maxWorkerId
	timestamp = (id >> timestampShift) + customEpoch

	return timestamp, workerId, sequence, nil
}

// GetOrderCreateTime 获取订单创建时间
func (g *OrderIdGenerator) GetOrderCreateTime(orderIdStr string) (time.Time, error) {
	timestamp, _, _, err := g.ParseOrderId(orderIdStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestamp/1000, (timestamp%1000)*1e6), nil
}

// waitNextMillis 等待下一毫秒
func (g *OrderIdGenerator) waitNextMillis(lastTimestamp int64) int64 {
	timestamp := getCurrentTimestamp()
	for timestamp <= lastTimestamp {
		time.Sleep(100 * time.Microsecond) // 短暂休眠
		timestamp = getCurrentTimestamp()
	}
	return timestamp
}

// getCurrentTimestamp 获取当前毫秒时间戳
func getCurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}

// 全局单例生成器
var (
	defaultGenerator *OrderIdGenerator
	once             sync.Once
)

// GetDefaultGenerator 获取默认生成器（单例）
func GetDefaultGenerator() *OrderIdGenerator {
	once.Do(func() {
		// 默认使用机器ID=1，实际使用时应该根据部署环境设置
		defaultGenerator, _ = NewOrderIdGenerator(1)
	})
	return defaultGenerator
}

// InitDefaultGenerator 初始化默认生成器
func InitDefaultGenerator(workerId int64) error {
	generator, err := NewOrderIdGenerator(workerId)
	if err != nil {
		return err
	}
	defaultGenerator = generator
	return nil
}

// GenerateOrderId 便捷方法：使用默认生成器生成订单ID
func GenerateOrderId() (string, error) {
	return GetDefaultGenerator().GenerateOrderId()
}

// GenerateId 便捷方法：使用默认生成器生成数字ID
func GenerateId() (int64, error) {
	return GetDefaultGenerator().GenerateId()
}

// ParseOrderId 便捷方法：解析订单ID
func ParseOrderId(orderIdStr string) (timestamp int64, workerId int64, sequence int64, err error) {
	return GetDefaultGenerator().ParseOrderId(orderIdStr)
}

// GetOrderCreateTime 便捷方法：获取订单创建时间
func GetOrderCreateTime(orderIdStr string) (time.Time, error) {
	return GetDefaultGenerator().GetOrderCreateTime(orderIdStr)
}
