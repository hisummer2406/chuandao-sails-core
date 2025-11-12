package logic

import (
	"chuandao-sails-core/app/api/delivery/internal/svc"
	"chuandao-sails-core/app/api/delivery/internal/types"
	"chuandao-sails-core/app/model"
	"chuandao-sails-core/app/pkg/engine/pricing"
	"chuandao-sails-core/common/snowflake"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询运费
func NewGetQuoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuoteLogic {
	return &GetQuoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetQuote API不做队列
func (l *GetQuoteLogic) GetQuote(req *types.GetQuotaReq) (resp *types.GetQuotaResp, err error) {
	// 1.幂等性检查，创建订单
	orderNo, err := l.ensureOrder(req)
	if err != nil {
		return nil, err
	}

	// 2.创建询价日志
	inquiryLog, err := l.createInquiryLog(orderNo, req)
	if err != nil {
		return nil, err
	}

	// 3.过滤禁用平台
	disablePlatforms := l.parseAblePlatforms(req.SelectDelivery)

	// 4.并发询价
	quoteReq := &pricing.QuoteRequest{
		OrderNo:          orderNo,
		FromLng:          req.FromAddress.Lng,
		FromLat:          req.FromAddress.Lat,
		FromAddress:      req.FromAddress.Detail,
		ToLng:            req.ToAddress.Lng,
		ToLat:            req.ToAddress.Lat,
		ToAddress:        req.ToAddress.Detail,
		GoodsType:        req.GoodsType,
		GoodsWeight:      req.GoodsWeight,
		SubscribeType:    req.SubscribeType,
		SubscribeTime:    req.SubscribeTime,
		DisablePlatforms: disablePlatforms,
	}

	quotes := l.svcCtx.PricingEngine.GetQuotes(l.ctx, quoteReq)

	// 5.保存询价明细
	if err := l.saveInquireDetails(inquiryLog.Id, quotes); err != nil {
		l.Errorf("[delivery-api] GetQuote save inquiry error: %v", err)
	}

	// 6.更新询价日志状态
	okCodes := l.extractOkPlatforms(quotes)
	l.updateInquiryLogStatus(inquiryLog.Id, okCodes)

	return &types.GetQuotaResp{
		orderNo,
		l.convertToQuotaList(quotes),
	}, nil
}

// ensureOrder 幂等性创建订单
func (l *GetQuoteLogic) ensureOrder(req *types.GetQuotaReq) (string, error) {
	// 1.查询订单是否存在
	existOrder, err := l.svcCtx.DispatchOrderModel.FindOneByUpstreamOrderId(l.ctx, req.UpstreamOrderId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return "", fmt.Errorf("查询订单失败: %v", err)
	}

	// 2.已存在直接返回
	if existOrder != nil {
		return existOrder.OrderNo, nil
	}

	// 3.生成新的订单号
	orderNo, err := snowflake.GenerateOrderNo()
	if err != nil {
		return "", fmt.Errorf("[delivery-api] GetQuoteLogic 'GenerateOrderNo' error:%v", err)
	}

	// 4.开启事务创建订单 + 日志
	err = l.svcCtx.DispatchOrderModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 4.1 构建订单对象
		fromAddress, _ := jsonx.Marshal(req.FromAddress)
		toAddress, _ := jsonx.Marshal(req.ToAddress)
		goodsDetail, _ := jsonx.Marshal(req.GoodInfo)

		order := &model.DispatchOrder{
			OrderNo:         orderNo,
			OriginOrderId:   req.OriginOrderId,
			UpstreamOrderId: req.UpstreamOrderId,
			UpstreamSource:  req.UpstreamSoucre,
			ShortNum:        req.ShortNum,
			Status:          "inquiring", // 询价中
			FromMobile:      req.FromAddress.Mobile,
			FromAddress:     string(fromAddress),
			ToMobile:        req.ToAddress.Mobile,
			ToAddress:       string(toAddress),
			Note:            req.Note,
			GoodsType:       req.GoodsType,
			GoodsDetail:     string(goodsDetail),
			SubscribeType:   req.SubscribeType,
			SubscribeTime:   req.SubscribeTime,
			SelectDelivery:  req.SelectDelivery,
		}

		// 4.2 插入订单主表
		_, err := l.svcCtx.DispatchOrderModel.InsertWithSessions(l.ctx, session, order)
		if err != nil {
			return fmt.Errorf("[delivery-api] GetQuote 'DispatchOrder Insert' error: %v", err)
		}

		// 4.3 插入订单状态日志
		statusLog := &model.DispatchOrderStatusLog{
			OrderNo:    orderNo,
			OldStatus:  "",
			NewStatus:  "inquiring",
			StatusDesc: "开始询价",
			Remark:     fmt.Sprintf("上游订单号：%s", req.UpstreamOrderId),
		}
		_, err = l.svcCtx.DispatchInquiryLogModel.InsertWithSession(ctx, session, statusLog)
		if err != nil {
			return fmt.Errorf("[delivery-api] GetQuote 'DispatchStatusLog Insert' error: %v", err)
		}
		return nil
	})
	if err != nil {
		return "", nil
	}

	return orderNo, nil
}

// createInquiryLog 创建询价日志
func (l *GetQuoteLogic) createInquiryLog(orderNo string, req *types.GetQuotaReq) (*model.DispatchInquiryLog, error) {
	inquiryLog := &model.DispatchInquiryLog{
		OrderNo:     orderNo,
		FromAddress: req.FromAddress.Detail,
		ToAddress:   req.ToAddress.Detail,
		Status:      "processing",
		GoodsType:   req.GoodsType,
	}

	result, err := l.svcCtx.DispatchInquiryLogModel.Insert(l.ctx, inquiryLog)
	if err != nil {
		return nil, fmt.Errorf("[delivery-api] GetQuote 'DispatchInquiryLog Insert' error: %v]")
	}

	inquiryLog.Id, _ = result.LastInsertId()
	return inquiryLog, nil
}

// parseDisablePlatforms 解析运力平台
func (l *GetQuoteLogic) parseAblePlatforms(delivery string) []string {
	if delivery == "" {
		// 切片是引用类型需要初始化
		return []string{}
	}
	return strings.Split(delivery, ",")
}

// saveInquireDetails 保存询价明细
func (l *GetQuoteLogic) saveInquireDetails(inquiryId int64, quotes []*pricing.QuoteResult) error {
	for _, quote := range quotes {
		detail := &model.DispatchInquiryDetail{
			InquiryId:      inquiryId,
			AccountId:      quote.AccountId,
			DeliveryCode:   quote.DeliveryCode,
			Price:          quote.Price,
			Distance:       quote.Distance,
			Duration:       quote.Duration,
			EstimateTime:   0, // 预计送达时间
			QuoteStatus:    boolToInt64(quote.Available),
			PriceToken:     "", // 价格令牌
			ResultResponse: quote.Reason,
		}
		_, err := l.svcCtx.DispatchInquiryDetailModel.Insert(l.ctx, detail)
		if err != nil {
			l.Errorf("[delivery-api] GetQuote 'DispatchInquiryDetail' error: %v", err)
		}
	}

	return nil
}

// updateInquiryLogStatus 更新询价日志状态
func (l *GetQuoteLogic) updateInquiryLogStatus(inquiryId int64, okCodes []string) {
	log, err := l.svcCtx.DispatchInquiryLogModel.FindOne(l.ctx, inquiryId)
	if err != nil {
		return
	}

	log.Status = "success"
	log.SuccessPlatforms = strings.Join(okCodes, ",")
	log.TotalDuration = 0 // TODO 计算总耗时

	l.svcCtx.DispatchInquiryLogModel.Update(l.ctx, log)
}

// extractOkPlatforms 提取成功的平台列表
func (l *GetQuoteLogic) extractOkPlatforms(quotes []*pricing.QuoteResult) []string {
	var codes []string
	for _, quote := range quotes {
		if quote.Available {
			codes = append(codes, quote.DeliveryCode)
		}
	}
	return codes
}

// convertToQuotaList logic响应格式
func (l *GetQuoteLogic) convertToQuotaList(quotes []*pricing.QuoteResult) []types.DeliveryQuota {
	var list []types.DeliveryQuota
	for _, quote := range quotes {
		list = append(list, types.DeliveryQuota{
			Delivery: types.Delivery{
				DeliveryCode: quote.DeliveryCode,
				DeliveryName: quote.DeliveryName,
			},
			Price:     quote.Price,
			Distance:  quote.Distance,
			Available: boolToInt64(quote.Available),
			Reason:    quote.Reason,
		})
	}
	return list
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
