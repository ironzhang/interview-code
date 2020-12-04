package access

import "net/http"

// 实现一个打印 http access log 的 Handler
// 会输出请求和响应日志
type AccessLogHandler struct {
	DisableBody bool         // 是否打印Body
	Handler     http.Handler // 实际的请求处理Handler
}

func (p *AccessLogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
