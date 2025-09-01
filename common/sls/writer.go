package sls

import (
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"strings"
	"time"
)

type Config struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	Project         string `json:"project"`
	Logstore        string `json:"logstore"`
	Topic           string `json:",optional"`
	Source          string `json:",optional"`
}

type Writer struct {
	producer *producer.Producer
	config   Config
}

// NewWriter 创建新的SLS日志写入器
func NewWriter(config Config) (*Writer, error) {

	// 使用新的CredentialsProducer
	credentialsProvider := sls.NewStaticCredentialsProvider(
		config.AccessKeyID,
		config.AccessKeySecret,
		config.SecurityToken, //临时访问令牌
	)

	// 创建Producer配置
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = config.Endpoint
	producerConfig.CredentialsProvider = credentialsProvider

	// 创建Producer实例
	producerInstance, err := producer.NewProducer(producerConfig)
	if err != nil {
		return nil, err
	}

	//启动 Producer
	producerInstance.Start()

	return &Writer{
		producer: producerInstance,
		config:   config,
	}, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	logContent := strings.TrimSpace(string(data))
	if logContent == "" {
		return len(data), nil
	}

	log := &sls.Log{
		Time: proto.Uint32(uint32(time.Now().Unix())),
		Contents: []*sls.LogContent{{
			Key:   proto.String("content"),
			Value: proto.String(logContent),
		}},
	}

	//使用SendLog 发送日志
	err := w.producer.SendLog(w.config.Project, w.config.Logstore, w.config.Topic, w.config.Source, log)
	if err != nil {
		return len(data), nil
	}
	return len(data), nil
}

// Alert 实现logx.Writer接口
func (w *Writer) Alert(v any) {
	w.writeLog("alert", v)
}

// Close 实现logx.Writer接口
func (w *Writer) Close() error {
	if w.producer != nil {
		w.producer.Close(5000) // 等待5秒安全关闭
	}
	return nil
}

// Debug 实现logx.Writer接口
func (w *Writer) Debug(v any, fields ...logx.LogField) {
	w.writeLog("debug", v, fields...)
}

// Error 实现logx.Writer接口
func (w *Writer) Error(v any, fields ...logx.LogField) {
	w.writeLog("error", v, fields...)
}

// Info 实现logx.Writer接口
func (w *Writer) Info(v any, fields ...logx.LogField) {
	w.writeLog("info", v, fields...)
}

// Severe 实现logx.Writer接口
func (w *Writer) Severe(v any) {
	w.writeLog("severe", v)
}

// Slow 实现logx.Writer接口
func (w *Writer) Slow(v any, fields ...logx.LogField) {
	w.writeLog("slow", v, fields...)
}

// Stack 实现logx.Writer接口
func (w *Writer) Stack(v any) {
	w.writeLog("stack", v)
}

// Stat 实现logx.Writer接口
func (w *Writer) Stat(v any, fields ...logx.LogField) {
	w.writeLog("stat", v, fields...)
}

// writeLog 写入日志
func (w *Writer) writeLog(level string, v any, fields ...logx.LogField) {
	contents := []*sls.LogContent{
		{
			Key:   proto.String("level"),
			Value: proto.String(level),
		}, {
			Key:   proto.String("message"),
			Value: proto.String(toString(v)),
		}, {
			Key:   proto.String("timestamp"),
			Value: proto.String(time.Now().Format("2006-01-02 15:04:05")),
		},
	}

	//添加自定义字段
	for _, field := range fields {
		contents = append(contents, &sls.LogContent{
			Key:   proto.String(field.Key),
			Value: proto.String(toString(field.Value)),
		})
	}

	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: contents,
	}

	// 发送日志，忽略错误
	w.producer.SendLog(w.config.Project, w.config.Logstore, w.config.Topic, w.config.Source, log)
}

// toString 将任意类型转换为字符串
func toString(v any) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%f", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case error:
		return val.Error()
	default:
		return fmt.Sprintf("%v", val)
	}
}
