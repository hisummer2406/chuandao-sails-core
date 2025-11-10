package middleware

import (
	"chuandao-sails-core/apps/platform-gateway/rpc/platformclient"
	CommonMiddleware "chuandao-sails-core/common/middleware"
	"context"
	"net/http"
)

type DeliveryConfigAdapter struct {
	platformRpc platformclient.Platform
}

// GetAppSecret 实现 common/sign 接口
func (a *DeliveryConfigAdapter) GetAppSecret(ctx context.Context, appId string) (secret string, status int, err error) {
	// 调用platform-gateway rpc
	resp, err := a.platformRpc.GetPlatformConfig(ctx, &platformclient.GetPlatformConfigReq{
		AppId: appId,
	})
	if err != nil {
		return "", 0, err
	}
	return resp.AppSecret, int(resp.Status), nil
}

type SignMiddleware struct {
	adapter *DeliveryConfigAdapter
}

func NewSignMiddleware(platformRpc platformclient.Platform) *SignMiddleware {
	return &SignMiddleware{
		adapter: &DeliveryConfigAdapter{
			platformRpc: platformRpc,
		},
	}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return CommonMiddleware.SignMiddlewareFunc(m.adapter)(next)
}
