package response

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"requestId"`
}

type BusinessError struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"requestId"`
}

func (b *BusinessError) Error() string {
	return b.Msg
}

const (
	SuccessCode      = 200 // 成功
	BusinessErrCode  = 400 // 业务错误
	UnauthorizedCode = 401 // 未授权
	ForbiddenCode    = 403 // 禁止访问
	NotFoundCode     = 404 // 资源不存在
	ParamErrCode     = 422 // 参数错误
	SystemErrCode    = 500 // 系统错误
)

// GetRequestID 获取请求ID，从go-zero 框架的TraceID 中获取
func GetRequestID(ctx context.Context) string {
	traceID := trace.TraceIDFromContext(ctx)
	if traceID == "" {
		traceID = "unknown"
	}

	return traceID
}

// NewBusinessErrorWithCtx 创建业务错误
func NewBusinessErrorWithCtx(ctx context.Context, code int, msg string) *BusinessError {
	return &BusinessError{
		Code:      code,
		Msg:       msg,
		Data:      make(map[string]interface{}),
		RequestID: GetRequestID(ctx),
	}
}

// NewParamErrorWithCtx 参数错误
func NewParamErrorWithCtx(ctx context.Context, msg string) *BusinessError {
	return &BusinessError{
		Code:      ParamErrCode,
		Msg:       msg,
		Data:      make(map[string]interface{}),
		RequestID: GetRequestID(ctx),
	}
}

// NewUnauthorizedErrorWithCtx 未授权错误
func NewUnauthorizedErrorWithCtx(ctx context.Context) *BusinessError {
	return &BusinessError{
		Code:      UnauthorizedCode,
		Msg:       "未授权访问",
		Data:      make(map[string]interface{}),
		RequestID: GetRequestID(ctx),
	}
}

// NewForbiddenErrorWithCtx 禁止访问错误
func NewForbiddenErrorWithCtx(ctx context.Context) *BusinessError {
	return &BusinessError{
		Code:      ForbiddenCode,
		Msg:       "禁止访问",
		Data:      make(map[string]interface{}),
		RequestID: GetRequestID(ctx),
	}
}

// NewNotFoundErrorWithCtx 资源不存在错误
func NewNotFoundErrorWithCtx(ctx context.Context) *BusinessError {
	return &BusinessError{
		Code:      NotFoundCode,
		Msg:       "资源不存在",
		Data:      make(map[string]interface{}),
		RequestID: GetRequestID(ctx),
	}
}

// NewSystemError 系统错误
func NewSystemError(requestID string) *BusinessError {
	return &BusinessError{
		Code:      SystemErrCode,
		Msg:       "系统错误",
		Data:      make(map[string]interface{}),
		RequestID: requestID,
	}
}

// HandleResponse 统一处理响应函数
func HandleResponse(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	if err != nil {
		//判断是否为业务错误
		if businessErr, ok := err.(*BusinessError); ok {
			//业务错误以及包含完整响应信息，直接返回
			httpx.OkJson(w, businessErr)
		} else {
			//系统错误，返回系统错误信息
			httpx.OkJson(w, &Response{
				Code:      SystemErrCode,
				Msg:       "系统内部错误",
				Data:      make(map[string]interface{}),
				RequestID: GetRequestID(r.Context()),
			})
		}
	} else {
		//成功返回
		httpx.OkJson(w, &Response{
			Code:      SuccessCode,
			Msg:       "success",
			Data:      data,
			RequestID: GetRequestID(r.Context()),
		})
	}
}
