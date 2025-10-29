package logic

import (
	"chuandao-sails-core/apps/delivery-dispatch/model"
	"chuandao-sails-core/apps/delivery-dispatch/pkg/engine/pricing"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"strings"

	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/svc"
	"chuandao-sails-core/apps/delivery-dispatch/api/v1/internal/types"

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

func (l *GetQuoteLogic) GetQuote(req *types.GetQuotaReq) (resp *types.GetQuotaResp, err error) {
	// 1.记录询价日志
	fromAddress, _ := jsonx.Marshal(req.FromAddress)
	toAddress, _ := jsonx.Marshal(req.ToAddress)

	inquiryLog := &model.DispatchInquiryLog{
		OrderNo:     req.OrderNO,
		FromAddress: string(fromAddress),
		ToAddress:   string(toAddress),
		Status:      "processing",
	}

	result, err := l.svcCtx.DispatchInquiryLogModel.Insert(l.ctx, inquiryLog)
	if err != nil {
		l.Errorf("GetQuote DispatchInquiryLogModel.Insert error: %v", err)
		return nil, err
	}
	inquiryId, _ := result.LastInsertId()

	// 2.询价请求
	disablePlatforms := []string{}
	if req.DisableDelivery != "" {
		disablePlatforms = strings.Split(req.DisableDelivery, ",")
	}

	quoteReq := &pricing.QuoteRequest{
		OrderNo:     req.OrderNO,
		FromLng:     req.FromAddress.Lng,
		FromLat:     req.FromAddress.Lat,
		FromAddress: req.FromAddress.Detail,
		ToLng:       req.ToAddress.Lng,
		ToLat:       req.ToAddress.Lat,
		ToAddress:   req.ToAddress.Detail,
	}

	// 3.并发询价

	// 4.保存询价明细

	// 5.更新询价日志状态

	// 6.响应
	return
}
