package pricing

import (
	"chuandao-sails-core/app/pkg/engine/platform"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type PricingEngine struct {
	platformManager *platform.PlatformManager
	timeout         time.Duration
}

// NewPricingEngine 创建询价引擎
func NewPricingEngine(platformManager *platform.PlatformManager, timeout time.Duration) *PricingEngine {
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	return &PricingEngine{
		platformManager: platformManager,
		timeout:         timeout,
	}
}

type (
	QuoteRequest struct {
		OrderNo          string   `json:"order_no"`
		FromLng          string   `json:"from_lng"` // 不用float64 影响建店精度
		FromLat          string   `json:"from_lat"`
		FromAddress      string   `json:"from_address"`
		ToLng            string   `json:"to_lng"`
		ToLat            string   `json:"to_lat"`
		ToAddress        string   `json:"to_address"`
		GoodsType        int64    `json:"goods_type"`
		GoodsWeight      float64  `json:"goods_weight"`
		SubscribeType    int64    `json:"subscribe_type"`    // 预约类型，默认实时订单
		SubscribeTime    string   `json:"subscribe_time"`    // 预约时间时间戳 pkg 转Time
		DisablePlatforms []string `json:"disable_platforms"` // 禁用询价平台列表
	}

	QuoteResult struct {
		OrderNo      string `json:"order_no"`
		DeliveryCode string `json:"delivery_code"`
		DeliveryName string `json:"delivery_name"`
		AccountId    int64  `json:"account_id"`
		AccountName  string `json:"account_name"`
		Price        int64  `json:"price"`         // 分
		Distance     int64  `json:"distance"`      // 米
		Available    bool   `json:"available"`     // 平台是否可用
		Reason       string `json:"reason"`        // 不可用时返回原因
		Duration     int64  `json:"duration"`      //
		ResponseTime int64  `json:"response_time"` // 相应时间(毫秒)
	}
)

// GetQuotes 并发询价（多平台多账号）
func (e *PricingEngine) GetQuotes(ctx context.Context, req *QuoteRequest) []*QuoteResult {
	// 1.获取所有可用平台的所有账号
	allAccounts := e.platformManager.GetAllAvailableAccounts(req.DisablePlatforms)
	if len(allAccounts) == 0 {
		logx.Errorf("No available delivery accounts")
		return nil
	}

	// 2.计算任务总数
	totalTasks := 0
	for _, account := range allAccounts {
		totalTasks += len(account)
	}

	// 3.创建结果通道
	results := make(chan *QuoteResult, totalTasks)

	// 4.创建带超时的上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	// 5.并发调用所有平台的所有账号
	var wg sync.WaitGroup
	for platformCode, accounts := range allAccounts {
		for _, account := range accounts {
			wg.Add(1)
			go e.queryPlatformAccount(timeoutCtx, &wg, platformCode, account, req, results)
		}
	}

	// 6.等待所有goroutine完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 7.收集结果
	return collectResults(results)
}

// queryPlatformAccount 查询单个平台账号(通道代码，不需要返回值)
func (e *PricingEngine) queryPlatformAccount(
	ctx context.Context,
	wg *sync.WaitGroup,
	deliveryCode string,
	account *platform.PlatformAccount,
	req *QuoteRequest,
	results chan<- *QuoteResult,
) {
	defer wg.Done()
	startTime := time.Now()

	// 1.获取平台适配器
	adapter, err := e.platformManager.GetAdapter(deliveryCode)
	if err != nil {
		logx.Errorf("Get adapter[%s] error: %v", deliveryCode, err)
		results <- &QuoteResult{
			OrderNo:      req.OrderNo,
			DeliveryCode: deliveryCode,
			AccountId:    account.ID,
			AccountName:  account.AccountName,
			Available:    false,
			Reason:       "平台适配器不存在",
		}
		return
	}

	// 2.构造平台询价请求
	platformReq := &platform.QuoteRequest{
		OrderNo:     req.OrderNo,
		FromLng:     req.FromLng,
		FromLat:     req.FromLat,
		FromAddress: req.FromAddress,
		ToLng:       req.ToLng,
		ToLat:       req.ToLat,
		ToAddress:   req.ToAddress,
		GoodsType:   req.GoodsType,
		GoodsWeight: req.GoodsWeight,
	}

	// 3.调用平台API
	resp, err := adapter.GetQuote(ctx, account, platformReq)
	responseTime := time.Since(startTime).Milliseconds()
	// 4.处理结果
	if err != nil {
		logx.Errorf("平台询价失败[%s-%s]: %v", deliveryCode, req.OrderNo, err)
		results <- &QuoteResult{
			OrderNo:      req.OrderNo,
			DeliveryCode: deliveryCode,
			DeliveryName: resp.DeliveryName,
			AccountId:    account.ID,
			AccountName:  account.AccountName,
			Available:    false,
			Reason:       err.Error(),
			ResponseTime: responseTime,
		}
		return
	}

	// 5.返回成功结果
	results <- &QuoteResult{
		OrderNo:      req.OrderNo,
		DeliveryCode: resp.DeliveryCode,
		DeliveryName: resp.DeliveryName,
		AccountId:    resp.AccountId,
		AccountName:  account.AccountName,
		Price:        resp.Price,
		Distance:     resp.Distance,
		Available:    resp.Available,
		Reason:       resp.Reason,
		ResponseTime: responseTime,
	}

}

// collectResults 接收所有结果
func collectResults(results <-chan *QuoteResult) []*QuoteResult {
	var allResults []*QuoteResult

	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}
