package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
)

// ResponseWriter wrapper to capture response data
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
	body       bytes.Buffer
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.size += len(data)
	// 记录响应数据
	if w.size < 3000 { // 限制缓存大小
		w.body.Write(data)
	}
	return w.ResponseWriter.Write(data)
}

// HTTPLogMiddleware 记录HTTP请求日志，简化版本
func HTTPLogMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// 获取trace信息
			ctx := r.Context()
			traceID := trace.TraceIDFromContext(ctx)
			spanID := trace.SpanIDFromContext(ctx)

			// 收集请求参数
			requestParams := collectRequestParams(r)

			// 包装ResponseWriter
			wrapped := &responseWriter{
				ResponseWriter: w,
				statusCode:     200,
			}

			// 执行处理器
			next(wrapped, r)

			// 计算耗时
			duration := time.Since(start)

			// go-zero格式日志消息
			message := fmt.Sprintf("[HTTP] %s - %s %s - %s - %s",
				getStatusCodeHighlight(wrapped.statusCode),
				getMethodHighlight(r.Method),
				r.URL.Path,
				getClientIP(r),
				getUserAgent(r),
			)

			// 只记录一条日志，包含所有必要信息
			logx.Infow(message,
				logx.Field("method", r.Method),
				logx.Field("path", r.URL.Path),
				logx.Field("status_code", wrapped.statusCode),
				logx.Field("duration", fmt.Sprintf("%.1fms", float64(duration.Nanoseconds())/1e6)),
				logx.Field("client_ip", getClientIP(r)),
				logx.Field("user_agent", getUserAgent(r)),
				logx.Field("request_params", requestParams),
				logx.Field("response_body", limitString(wrapped.body.String(), 1500)),
				logx.Field("response_size", wrapped.size),
				logx.Field("trace", traceID),
				logx.Field("span", spanID),
				logx.Field("caller", getCaller()),
			)
		}
	}
}

// getCaller 自动获取调用者信息
func getCaller() string {
	// 跳过runtime调用栈，找到真正的调用者
	// 0: getCaller
	// 1: HTTPLogMiddleware的匿名函数
	// 2: 中间件包装函数
	// 3: go-zero的路由处理器
	// 4: 实际的处理函数
	for skip := 4; skip <= 10; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		funcName := fn.Name()

		// 跳过go-zero内部和中间件的调用
		if strings.Contains(funcName, "github.com/zeromicro/go-zero") ||
			strings.Contains(funcName, "middleware.") ||
			strings.Contains(file, "middleware/") {
			continue
		}

		// 简化文件路径，只保留关键部分
		if idx := strings.LastIndex(file, "/"); idx >= 0 {
			file = file[idx+1:]
		}

		return fmt.Sprintf("%s:%d", file, line)
	}

	// 如果找不到合适的调用者，返回默认值
	return "unknown:0"
}

// collectRequestParams 收集所有请求参数到一个字段
func collectRequestParams(r *http.Request) string {
	params := make(map[string]interface{})

	// 查询参数
	if len(r.URL.RawQuery) > 0 {
		queryParams := r.URL.Query()
		if len(queryParams) > 0 {
			params["query"] = flattenParams(queryParams)
		}
	}

	// 请求体
	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil && len(bodyBytes) > 0 {
			bodyStr := string(bodyBytes)
			// 重新设置body
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// 尝试解析JSON，失败则直接存储字符串
			var jsonData interface{}
			if err := json.Unmarshal(bodyBytes, &jsonData); err == nil {
				params["body"] = jsonData
			} else {
				params["body"] = limitString(bodyStr, 1000)
			}
		}
	}

	// 表单参数
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		contentType := r.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			r.ParseForm()
			if len(r.PostForm) > 0 {
				params["form"] = flattenParams(r.PostForm)
			}
		}
	}

	// 重要的请求头
	headers := make(map[string]string)
	if auth := r.Header.Get("Authorization"); auth != "" {
		headers["authorization"] = "***" // 隐藏敏感信息
	}
	if contentType := r.Header.Get("Content-Type"); contentType != "" {
		headers["content_type"] = contentType
	}
	if len(headers) > 0 {
		params["headers"] = headers
	}

	// 如果没有任何参数，返回空字符串
	if len(params) == 0 {
		return ""
	}

	data, _ := json.Marshal(params)
	return string(data)
}

// flattenParams 扁平化URL参数
func flattenParams(params url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range params {
		if len(v) == 1 {
			result[k] = v[0]
		} else {
			result[k] = v
		}
	}
	return result
}

// getStatusCodeHighlight 根据状态码返回带背景色和加粗的字符串
func getStatusCodeHighlight(code int) string {
	switch {
	case code >= 200 && code < 300:
		return fmt.Sprintf("\033[1;97;42m %d \033[0m", code)
	case code >= 300 && code < 400:
		return fmt.Sprintf("\033[1;30;43m %d \033[0m", code)
	case code >= 400 && code < 500:
		return fmt.Sprintf("\033[1;97;41m %d \033[0m", code)
	case code >= 500:
		return fmt.Sprintf("\033[1;97;45m %d \033[0m", code)
	default:
		return fmt.Sprintf("%d", code)
	}
}

// getMethodHighlight 根据HTTP方法返回带背景色和加粗的字符串
func getMethodHighlight(method string) string {
	switch method {
	case "GET":
		return fmt.Sprintf("\033[1;97;46m %s \033[0m", method)
	case "POST":
		return fmt.Sprintf("\033[1;97;42m %s \033[0m", method)
	case "PUT":
		return fmt.Sprintf("\033[1;30;43m %s \033[0m", method)
	case "DELETE":
		return fmt.Sprintf("\033[1;97;41m %s \033[0m", method)
	case "PATCH":
		return fmt.Sprintf("\033[1;97;45m %s \033[0m", method)
	default:
		return method
	}
}

// limitString 限制字符串长度
func limitString(s string, limit int) string {
	if len(s) <= limit {
		return s
	}
	return s[:limit] + "..."
}

// getClientIP 获取客户端真实IP
func getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}

// getUserAgent 获取User-Agent，截断过长的
func getUserAgent(r *http.Request) string {
	ua := r.UserAgent()
	if len(ua) > 60 {
		return ua[:60] + "..."
	}
	return ua
}
