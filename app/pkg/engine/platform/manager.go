package platform

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ============ 核心数据结构 ============

type (

	// PlatformConfig 平台配置（映射 delivery_platform_config表）
	PlatformConfig struct {
		ID           int64  `db:"id"`
		PlatformCode string `db:"platform_code"` // UU/DD/SF
		PlatformName string `db:"platform_name"`
		Status       int    `db:"status"`  // 1-启用 0-禁用
		ApiUrl       string `db:"api_url"` // API地址
	}

	// PlatformAccount 账号配置（映射 delivery_platform_account 表）
	PlatformAccount struct {
		ID           int64  `db:"id"`
		PlatformCode string `db:"platform_code"` // 关联平台
		AccountCode  string `db:"account_code"`
		AccountName  string `db:"account_name"`
		AppId        string `db:"app_id"`
		AppKey       string `db:"app_key"`
		AppSecret    string `db:"app_secret"`
		Status       int    `db:"status"`       // 1-启用 0-禁用
		ExtraConfig  string `db:"extra_config"` // JSON 扩展配置
	}
)

// ============ 适配器接口 ============

// PlatformAdapter 平台适配器统一接口
type PlatformAdapter interface {
	// GetQuote 询价
	GetQuote(ctx context.Context, account *PlatformAccount, req *QuoteRequest) (*QuoteResponse, error)
	// DispatchOrder 发单
	DispatchOrder(ctx context.Context, account *PlatformAccount, req *DispatchOrderRequest) (*DispatchOrderResponse, error)
	// GetCancelFee 获取违约金
	GetCancelFee(ctx context.Context, account *PlatformAccount, PlatformOrderNo string) error
	// CancelOrder 取消订单(TODO 取消原因写死)
	CancelOrder(ctx context.Context, account *PlatformAccount, PlatformOrderNo string) error
	// GetDriverLocation 获取骑手位置
	GetDriverLocation(ctx context.Context, account *PlatformAccount, PlatformOrderNo string) (string, error)
	// AddTip 加小费
	AddTip(ctx context.Context, account *PlatformAccount, PlatformOrderNo string, amount int64) error
	// QueryOrder 查询订单
	QueryOrder(ctx context.Context, account *PlatformAccount, PlatformOrderNo string) (*OrderDetail, error)
}

// ============ 请求响应结构 ============

// 询价结构
type (
	QuoteRequest struct {
		OrderNo     string  `json:"order_no"`
		CityCode    string  `json:"city_code"`
		FromLng     string  `json:"from_lng"` // 发货地址
		FromLat     string  `json:"from_lat"`
		FromAddress string  `json:"from_address"`
		ToLng       string  `json:"to_lng"` // 收货地址
		ToLat       string  `json:"to_lat"`
		ToAddress   string  `json:"to_address"`
		GoodsType   int64   `json:"goods_type"` // TODO 注意兼容
		GoodsWeight float64 `json:"goods_weight"`
	}

	QuoteResponse struct {
		OrderNo      string `json:"order_no"`
		DeliveryCode string `json:"delivery_code"` // 平台
		DeliveryName string `json:"delivery_name"`
		AccountId    int64  `json:"account_id"` // 账号ID
		Price        int64  `json:"price"`      // 报价
		Distance     int64  `json:"distance"`   // 距离
		Available    bool   `json:"available"`  // 平台是否可用
		Reason       string `json:"reason"`     // 不可用时返回原因
		Duration     int64  `json:"duration"`   //
	}
)

// 发单结构
type (
	DispatchOrderRequest struct {
		OrderNo     string  `json:"order_no"`
		FromName    string  `json:"from_name"` // 发货信息
		FromMobile  string  `json:"from_mobile"`
		FromLng     string  `json:"from_lng"`
		FromLat     string  `json:"from_lat"`
		FromAddress string  `json:"from_address"`
		ToName      string  `json:"to_name"` // 收货信息
		ToMobile    string  `json:"to_mobile"`
		ToLng       string  `json:"to_lng"`
		ToLat       string  `json:"to_lat"`
		ToAddress   string  `json:"to_address"`
		GoodsName   string  `json:"goods_name"` // 物品信息
		GoodsType   int     `json:"goods_type"`
		GoodsWeight float64 `json:"goods_weight"`
		Remark      string  `json:"remark"`       // 备注信息
		CallbackUrl string  `json:"callback_url"` // 订单状态回调地址
	}

	DispatchOrderResponse struct {
		PlatformOrderNo string `json:"platform_order_no"`
		Status          int    `json:"status"`
		EstimatedTime   int    `json:"estimated_time"` // 预估时间
		Fee             int64  `json:"fee"`            // 发单价格
	}
)

// 查询订单
type OrderDetail struct {
	PlatformOrderNo string    `db:"platform_order_no"`
	Status          int       `db:"status"`
	StatusText      string    `db:"status_text"`
	DriverName      string    `db:"driver_name"`
	DriverMobile    string    `db:"driver_mobile"`
	CreateTime      time.Time `db:"create_time"`  // 发单时间
	AcceptTime      time.Time `db:"accept_time"`  // 接单时间
	PickupTime      time.Time `db:"pickup_time"`  // 取货时间
	DeliverTime     time.Time `db:"deliver_time"` // 送达时间
}

// 骑手位置
type DriverLocation struct {
	DriverName   string    `db:"driver_name"`
	DriverMobile string    `db:"driver_mobile"`
	LastLng      float64   `db:"last_lng"`
	LastLat      float64   `db:"last_lat"`
	UpdateTime   time.Time `db:"update_time"`
}

// PlatformQuoteResult 平台询价结果
type PlatformQuoteResult struct {
	PlatformCode string
	Account      *PlatformAccount
	Quote        *QuoteResponse
	Error        error
}

// ============ 平台管理器 ============

// PlatformManager 平台管理器（负责平台和账号的管理）
type PlatformManager struct {
	configs  map[string]*PlatformConfig
	accounts map[string][]*PlatformAccount
	adapters map[string]PlatformAdapter
	mu       sync.RWMutex // 读写互斥锁，读多写少的场景
}

// NewPlatformManager 创建平台管理器
func NewPlatformManager() *PlatformManager {
	return &PlatformManager{
		configs:  make(map[string]*PlatformConfig),
		accounts: make(map[string][]*PlatformAccount),
		adapters: make(map[string]PlatformAdapter),
	}
}

// RegisterAdapter 注册平台适配器
func (m *PlatformManager) RegisterAdapter(platformCode string, adapter PlatformAdapter) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.adapters[platformCode] = adapter
}

// LoadConfig 加载平台配置
func (m *PlatformManager) LoadConfig(configs []*PlatformConfig) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, config := range configs {
		if config.Status == 1 {
			m.configs[config.PlatformCode] = config
		}
	}
}

// LoadAccounts 加载平台账号
func (m *PlatformManager) LoadAccounts(accounts []*PlatformAccount) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 清空旧数据，(make 返回Type本身，用于slice map channel)
	m.accounts = make(map[string][]*PlatformAccount)

	// 按平台分组
	for _, account := range accounts {
		if account.Status == 1 { // 只加载启用的账号
			m.accounts[account.PlatformCode] = append(m.accounts[account.PlatformCode], account)
		}
	}
}

// GetAvailablePlatforms 获取可用的平台列表（排除禁用的）
func (m *PlatformManager) GetAvailablePlatforms(disableCodes []string) []*PlatformConfig {
	m.mu.RLock()
	defer m.mu.RUnlock()

	disableMap := make(map[string]bool)
	for _, code := range disableCodes {
		disableMap[code] = true
	}

	var platformCodes []*PlatformConfig
	for code, config := range m.configs {
		if !disableMap[code] {
			platformCodes = append(platformCodes, config)
		}
	}

	return platformCodes
}

// =========== 不选择账号，从发单策略中获取 ==========

// GetAdapter 获取平台适配器
func (m *PlatformManager) GetAdapter(platformCode string) (PlatformAdapter, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	adapter, ok := m.adapters[platformCode]
	if !ok {
		return nil, fmt.Errorf("unknown platform code: %s", platformCode)
	}

	return adapter, nil
}

// GetAllAvailableAccounts 获取所有可用的平台账号（用于并发询价）
func (m *PlatformManager) GetAllAvailableAccounts(disableCodes []string) map[string][]*PlatformAccount {
	m.mu.RLock()
	defer m.mu.RUnlock()

	disableMap := make(map[string]bool)
	for _, code := range disableCodes {
		disableMap[code] = true
	}

	// 一个平台多账号 map[string][]*PlatformAccount
	result := make(map[string][]*PlatformAccount)
	for platformCode, accounts := range m.accounts {
		if !disableMap[platformCode] {
			continue
		}

		if _, ok := m.configs[platformCode]; !ok {
			continue
		}
		result[platformCode] = accounts
	}

	return result
}

// GetAccountByPlatformCode 获取平台的所有账号
func (m *PlatformManager) GetAccountsByPlatformCode(platformCode string) []*PlatformAccount {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.accounts[platformCode]
}

// GetAccountID 根据ID获取账号
func (m *PlatformManager) GetAccountID(accountId int64) (*PlatformAccount, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, accountList := range m.accounts {
		for _, account := range accountList {
			if account.ID == accountId {
				return account, nil
			}
		}
	}

	return nil, fmt.Errorf("unknown account id: %d", accountId)
}
