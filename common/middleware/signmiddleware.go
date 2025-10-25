package middleware

import (
	"bytes"
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

// SignRequest 签名请求参数
type SignRequest struct {
	AppId     string `json:"appId"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Sign      string `json:"sign"`
}

// PlatformConfigGetter 平台配置获取接口
type PlatformConfigGetter interface {
	GetAppSecret(ctx context.Context, appId string) (secret string, status int, err error)
}

// 通过接口依赖注入，任何域都可以使用
// SignMiddlewareFunc 通用签名验证中间件函数
func SignMiddlewareFunc(configGetter PlatformConfigGetter) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 1.读取请求体
			body, err := io.ReadAll(r.Body)
			if err != nil {
				response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), "读取请求体失败"))
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(body))

			// 2.解析签名参数
			var req SignRequest
			if err := jsonx.Unmarshal(body, &req); err != nil {
				response.HandleResponse(w, r, nil, response.NewParamErrorWithCtx(r.Context(), "缺少必要的签名参数"))
				return
			}

			// 3.获取平台配置
			appSecret, status, err := configGetter.GetAppSecret(r.Context(), req.AppId)
			if err != nil {
				response.HandleResponse(w, r, nil, response.NewUnauthorizedErrorWithCtx(r.Context()))
				return
			}

			// 4.检查平台状态
			if status != 1 {
				response.HandleResponse(w, r, nil, response.NewForbiddenErrorWithCtx(r.Context()))
				return
			}

			// 5.验证签名
			signStr := req.Data + appSecret + req.Timestamp
			expectedSign := Md5HashUpper(signStr)

			if !strings.EqualFold(req.Sign, expectedSign) {
				logc.Error(r.Context(), "签名验证失败",
					logc.Field("appId", req.AppId),
					logc.Field("expectedSign", expectedSign),
					logc.Field("reqSign", req.Sign),
				)
			}
			response.HandleResponse(w, r, nil, response.NewUnauthorizedErrorWithCtx(r.Context()))
			return

			// 6.继续处理
			next(w, r)
		}
	}
}

// Md5HashUpper 计算md5签名并转为16进制32位大写字符串
func Md5HashUpper(data string) string {
	hash := md5.Sum([]byte(data))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
