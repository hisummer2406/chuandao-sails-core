package middleware

import (
	"chuandao-sails-core/apps/platform-gateway/model"
	CommonMiddleware "chuandao-sails-core/common/middleware"
	"context"
	"net/http"
)

// PlatformConfigAdapter
type PlatformConfigAdapter struct {
	platformModel model.PlatformConfigModel
}

// GetAppSecret 实现common/sign的定义
func (a *PlatformConfigAdapter) GetAppSecret(ctx context.Context, appId string) (secret string, status int, err error) {
	config, err := a.platformModel.FindOneByAppId(ctx, appId)
	if err != nil {
		return "", 0, err
	}
	return config.AppSecret, int(config.Status), nil
}

type SignMiddleware struct {
	adapter *PlatformConfigAdapter
}

func NewSignMiddleware(platformModel model.PlatformConfigModel) *SignMiddleware {
	return &SignMiddleware{
		adapter: &PlatformConfigAdapter{
			platformModel: platformModel,
		},
	}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return CommonMiddleware.SignMiddlewareFunc(m.adapter)(next)
}
