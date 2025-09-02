package middleware

import (
	CommonMiddleware "chuandao-sails-core/common/middleware"
	"net/http"
)

type HTTPLogMiddleware struct {
}

func NewHTTPLogMiddleware() *HTTPLogMiddleware {
	return &HTTPLogMiddleware{}
}

func (m *HTTPLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	// 直接使用 common/middleware 中的 HTTPLogMiddleware 函数
	return CommonMiddleware.HTTPLogMiddleware()(next)
}
