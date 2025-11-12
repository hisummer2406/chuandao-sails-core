package middleware

import (
	"chuandao-sails-core/app/model"
	CommonMiddleware "chuandao-sails-core/common/middleware"
	"context"
	"net/http"
)

type DeliveryConfigAdapter struct {
	platformModel model.GatewayPlatformConfigModel
}

// GetAppSecret 实现 common/sign 接口
func (a *DeliveryConfigAdapter) GetAppSecret(ctx context.Context, appId string) (secret string, status int, err error) {
	// 调用platform-gateway rpc
	resp, err := a.platformModel.FindOneByAppId(ctx, appId)
	if err != nil {
		return "", 0, err
	}
	return resp.AppSecret, int(resp.Status), nil
}

type SignMiddleware struct {
	adapter *DeliveryConfigAdapter
}

func NewSignMiddleware(platformModel model.GatewayPlatformConfigModel) *SignMiddleware {
	return &SignMiddleware{
		adapter: &DeliveryConfigAdapter{
			platformModel: platformModel,
		},
	}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return CommonMiddleware.SignMiddlewareFunc(m.adapter)(next)
}
