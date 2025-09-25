package snowflake

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Example_基本使用
func Example_basicUsage() {
	// 1. 创建生成器（机器ID=1）
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		panic(err)
	}

	// 2. 生成订单ID
	orderId, err := generator.GenerateOrderId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("生成的订单ID: %s\n", orderId)

	// 3. 生成数字ID
	numId, err := generator.GenerateId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("生成的数字ID: %d\n", numId)

	// 4. 自定义前缀
	customOrderId, err := generator.GenerateOrderIdWithPrefix("MEAL")
	if err != nil {
		panic(err)
	}
	fmt.Printf("自定义前缀订单ID: %s\n", customOrderId)
}

// Example_全局单例使用
func Example_singletonUsage() {
	// 1. 初始化默认生成器
	err := InitDefaultGenerator(2) // 机器ID=2
	if err != nil {
		panic(err)
	}

	// 2. 直接使用便捷方法
	orderId, err := GenerateOrderId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("单例生成的订单ID: %s\n", orderId)

	// 3. 解析订单ID
	timestamp, workerId, sequence, err := ParseOrderId(orderId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("时间戳: %d, 机器ID: %d, 序列号: %d\n", timestamp, workerId, sequence)

	// 4. 获取创建时间
	createTime, err := GetOrderCreateTime(orderId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("订单创建时间: %s\n", createTime.Format("2006-01-02 15:04:05"))
}

// Example_批量生成
func Example_batchGenerate() {
	generator, _ := NewOrderIdGenerator(3)

	// 批量生成10个订单ID
	orderIds, err := generator.BatchGenerate(10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("批量生成的订单ID:\n")
	for i, id := range orderIds {
		fmt.Printf("%d: %s\n", i+1, id)
	}
}

// TestGenerateOrderId_ShowOutput 查看订单ID生成结果
func TestGenerateOrderId_ShowOutput(t *testing.T) {
	// 初始化生成器
	err := InitDefaultGenerator(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("\n=== 订单ID生成结果测试 ===")

	// 生成单个订单ID
	orderId, err := GenerateOrderId()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("生成的订单ID: %s\n", orderId)
	fmt.Printf("ID长度: %d 字符\n", len(orderId))
	fmt.Printf("前缀: %s\n", orderId[:3])
	fmt.Printf("数字部分: %s\n", orderId[3:])

	// 解析订单ID获取详细信息
	timestamp, workerId, sequence, err := ParseOrderId(orderId)
	if err != nil {
		t.Fatal(err)
	}

	createTime, _ := GetOrderCreateTime(orderId)

	fmt.Println("\n--- ID解析结果 ---")
	fmt.Printf("时间戳: %d\n", timestamp)
	fmt.Printf("机器ID: %d\n", workerId)
	fmt.Printf("序列号: %d\n", sequence)
	fmt.Printf("创建时间: %s\n", createTime.Format("2006-01-02 15:04:05.000"))

	// 验证基本属性
	if len(orderId) != 22 {
		t.Errorf("期望ID长度22，实际%d", len(orderId))
	}
	if orderId[:3] != "ORD" {
		t.Errorf("期望前缀ORD，实际%s", orderId[:3])
	}
	if workerId != 1 {
		t.Errorf("期望机器ID为1，实际%d", workerId)
	}
}

// Test_并发安全性测试
func Test_ConcurrentSafety(t *testing.T) {
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		t.Fatal(err)
	}

	const goroutines = 10
	const idsPerGoroutine = 1000

	var wg sync.WaitGroup
	idChan := make(chan string, goroutines*idsPerGoroutine)

	// 启动多个goroutine并发生成ID
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(workerNum int) {
			defer wg.Done()
			for j := 0; j < idsPerGoroutine; j++ {
				orderId, err := generator.GenerateOrderId()
				if err != nil {
					t.Errorf("Worker %d: %v", workerNum, err)
					return
				}
				idChan <- orderId
			}
		}(i)
	}

	wg.Wait()
	close(idChan)

	// 检查ID唯一性
	idSet := make(map[string]bool)
	duplicates := 0

	for orderId := range idChan {
		if idSet[orderId] {
			duplicates++
			t.Errorf("发现重复ID: %s", orderId)
		}
		idSet[orderId] = true
	}

	expectedCount := goroutines * idsPerGoroutine
	actualCount := len(idSet)

	if actualCount != expectedCount {
		t.Errorf("期望生成 %d 个唯一ID，实际生成 %d 个", expectedCount, actualCount)
	}

	if duplicates > 0 {
		t.Errorf("发现 %d 个重复ID", duplicates)
	}

	t.Logf("并发测试通过：生成了 %d 个唯一ID", actualCount)
}

// Test_性能测试
func Test_Performance(t *testing.T) {
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		t.Fatal(err)
	}

	const count = 100000
	start := time.Now()

	for i := 0; i < count; i++ {
		_, err := generator.GenerateId()
		if err != nil {
			t.Fatal(err)
		}
	}

	duration := time.Since(start)
	qps := float64(count) / duration.Seconds()

	t.Logf("生成 %d 个ID 耗时: %v", count, duration)
	t.Logf("性能: %.2f IDs/秒", qps)

	// 性能要求：至少10万QPS
	if qps < 100000 {
		t.Errorf("性能不达标，期望至少100000 QPS，实际 %.2f QPS", qps)
	}
}

// Test_ID解析测试
func Test_ParseOrderId(t *testing.T) {
	generator, err := NewOrderIdGenerator(123)
	if err != nil {
		t.Fatal(err)
	}

	// 生成订单ID
	orderId, err := generator.GenerateOrderId()
	if err != nil {
		t.Fatal(err)
	}

	// 解析订单ID
	timestamp, workerId, sequence, err := generator.ParseOrderId(orderId)
	if err != nil {
		t.Fatal(err)
	}

	// 验证机器ID
	if workerId != 123 {
		t.Errorf("机器ID不匹配，期望 123，实际 %d", workerId)
	}

	// 验证时间戳合理性（应该是最近的时间）
	now := time.Now().UnixMilli()
	if timestamp < now-1000 || timestamp > now+1000 {
		t.Errorf("时间戳不合理，生成时间: %d, 当前时间: %d", timestamp, now)
	}

	// 验证序列号范围
	if sequence < 0 || sequence > maxSequence {
		t.Errorf("序列号超出范围，序列号: %d", sequence)
	}

	t.Logf("解析成功 - 订单ID: %s, 时间戳: %d, 机器ID: %d, 序列号: %d",
		orderId, timestamp, workerId, sequence)
}

// Test_时钟回拨测试
func Test_ClockBackwards(t *testing.T) {
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		t.Fatal(err)
	}

	// 生成一个ID
	_, err = generator.GenerateId()
	if err != nil {
		t.Fatal(err)
	}

	// 模拟时钟回拨：手动设置一个未来的时间戳
	generator.lastTimestamp = time.Now().UnixMilli() + 5000

	// 尝试再次生成，应该返回错误
	_, err = generator.GenerateId()
	if err == nil {
		t.Error("期望时钟回拨错误，但没有返回错误")
	}

	t.Logf("时钟回拨检测正常: %v", err)
}

// Benchmark_生成订单ID性能基准测试
func Benchmark_GenerateOrderId(b *testing.B) {
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := generator.GenerateOrderId()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// Benchmark_生成数字ID性能基准测试
func Benchmark_GenerateId(b *testing.B) {
	generator, err := NewOrderIdGenerator(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := generator.GenerateId()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Example_inBusinessCode 示例：在业务代码中的使用
func Example_inBusinessCode() {
	// 在应用启动时初始化
	// workerId 可以从配置文件、环境变量或服务发现中获取
	workerId := getWorkerIdFromConfig() // 自定义函数
	err := InitDefaultGenerator(workerId)
	if err != nil {
		panic(fmt.Sprintf("初始化订单ID生成器失败: %v", err))
	}

	// 在业务逻辑中使用
	orderId, err := GenerateOrderId()
	if err != nil {
		// 处理错误
		return
	}

	fmt.Printf("新订单ID: %s\n", orderId)
}

// 模拟配置获取函数
func getWorkerIdFromConfig() int64 {
	// 实际项目中可能从以下来源获取：
	// 1. 配置文件
	// 2. 环境变量
	// 3. 服务注册中心
	// 4. K8S Pod 名称解析
	// 5. MAC 地址 hash
	return 1
}
