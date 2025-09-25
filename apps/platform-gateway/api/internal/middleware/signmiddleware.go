package middleware

import (
	"bytes"
	"chuandao-sails-core/apps/platform-gateway/api/types"
	"chuandao-sails-core/apps/platform-gateway/model"
	"chuandao-sails-core/common/response"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"net/http"
	"strings"
)

type SignMiddleware struct {
	platformModel model.PlatformConfigModel
}

func NewSignMiddleware(pm model.PlatformConfigModel) *SignMiddleware {
	return &SignMiddleware{
		platformModel: pm,
	}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), "读取请求体失败"))
			return
		}
		// 重新设置请求体，确保Handler可以正常解析参数
		r.Body = io.NopCloser(bytes.NewReader(body))

		//解析参数
		var req types.SignType
		if err := jsonx.Unmarshal(body, &req); err != nil {
			response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), "缺少必要的签名参数"))
			return
		}
		//获取配置
		config, err := m.platformModel.FindOneByAppId(context.Background(), req.AppId)
		if err != nil {
			response.HandleResponse(w, r, nil, response.NewUnauthorizedErrorWithCtx(r.Context()))
			return
		}

		if config.Status != 1 {
			response.HandleResponse(w, r, nil, response.NewForbiddenErrorWithCtx(r.Context()))
			return
		}

		//验证签名:data + secret + timestamp
		signStr := req.Data + config.AppSecret + req.Timestamp
		expectedSign := m.md5Hash(signStr)

		if !strings.EqualFold(req.Sign, expectedSign) {
			logc.Error(r.Context(), "签名验证失败",
				logc.Field("appId", req.AppId),
				logc.Field("expectedSign", expectedSign),
				logc.Field("reqSign", req.Sign),
			)
			response.HandleResponse(w, r, nil, response.NewUnauthorizedErrorWithCtx(r.Context()))
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}

// md5Hash 计算MD5哈希值并转为16进制32位大写字符串
func (m *SignMiddleware) md5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
