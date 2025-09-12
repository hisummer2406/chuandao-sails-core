package rocketmq

import "time"

type Config struct {
	//连接配置
	Endpoint      string         `json:"endpoint"`
	AccessKey     string         `json:"access_key"`
	SecretKey     string         `json:"secret_key"`
	NameSpace     string         `json:"namespace"`      //实例id
	ConsumerGroup string         `json:"consumer_group"` //消费组
	Producer      ProducerConfig `json:"producer"`       //生产者配置
	Consumer      ConsumerConfig `json:"consumer"`       //消费者配置
}

type ProducerConfig struct {
	Topics  []string      `json:"topics"`  //支持的主题列表
	Timeout time.Duration `json:"timeout"` //发送超时
	Retries int           `json:"retries"` //重试次数
}

type ConsumerConfig struct {
	Group             string        `json:"group"`
	AwaitDuration     time.Duration `json:"await_duration"`     //等待时长
	MaxMessageNum     int32         `json:"max_message_num"`    //批量拉取消息数
	InvisibleDuration time.Duration `json:"invisible_duration"` //不可见时长
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Producer: ProducerConfig{
			Topics:  []string{"demo-mp-helloworld"},
			Timeout: 3 * time.Second,
			Retries: 3,
		},
		Consumer: ConsumerConfig{
			Group:             "demo-mp-gr",
			AwaitDuration:     15 * time.Second,
			MaxMessageNum:     32,
			InvisibleDuration: 20 * time.Second,
		},
	}
}
