package snowflake

// Config 机器ID
type Config struct {
	WorkerId int64 `json:"worker_id"` // 机器ID，范围0-1023
}
