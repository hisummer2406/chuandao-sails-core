package model

import (
	"encoding/json"
	"time"
)

// OrderMessage 订单消息 - 放在业务代码中
type OrderMessage struct {
	OrderID     string                 `json:"order_id"`
	UserID      string                 `json:"user_id"`
	Amount      float64                `json:"amount"`
	Status      string                 `json:"status"`
	ProductInfo map[string]interface{} `json:"product_info,omitempty"`
	CreateTime  time.Time              `json:"create_time"`
	UpdateTime  time.Time              `json:"update_time"`
}

func (o *OrderMessage) GetTopic() string         { return "order_topic" }
func (o *OrderMessage) GetTag() string           { return o.Status }
func (o *OrderMessage) GetKeys() []string        { return []string{o.OrderID} }
func (o *OrderMessage) ToBytes() ([]byte, error) { return json.Marshal(o) }

// PaymentMessage 支付消息 - 放在业务代码中
type PaymentMessage struct {
	PaymentID   string    `json:"payment_id"`
	OrderID     string    `json:"order_id"`
	UserID      string    `json:"user_id"`
	Amount      float64   `json:"amount"`
	Method      string    `json:"method"`
	Status      string    `json:"status"`
	ProcessTime time.Time `json:"process_time"`
}

func (p *PaymentMessage) GetTopic() string         { return "payment_topic" }
func (p *PaymentMessage) GetTag() string           { return p.Status }
func (p *PaymentMessage) GetKeys() []string        { return []string{p.PaymentID, p.OrderID} }
func (p *PaymentMessage) ToBytes() ([]byte, error) { return json.Marshal(p) }

// InventoryMessage 库存消息
type InventoryMessage struct {
	ProductID   string `json:"product_id"`
	SKU         string `json:"sku"`
	Quantity    int    `json:"quantity"`
	Operation   string `json:"operation"` // LOCK, UNLOCK, DEDUCT
	OrderID     string `json:"order_id,omitempty"`
	WarehouseID string `json:"warehouse_id,omitempty"`
}

func (i *InventoryMessage) GetTopic() string         { return "inventory_topic" }
func (i *InventoryMessage) GetTag() string           { return i.Operation }
func (i *InventoryMessage) GetKeys() []string        { return []string{i.ProductID, i.OrderID} }
func (i *InventoryMessage) ToBytes() ([]byte, error) { return json.Marshal(i) }
