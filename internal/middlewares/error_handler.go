package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ErrorInfo 错误信息结构
type ErrorInfo struct {
	Code     int
	Message  string
	Detail   string // 详细信息（可选）
	CustomJS string // 自定义 JavaScript 代码（可选）
}

// 上下文中存储自定义错误信息的 key
const ctxKeyErrorInfo = "ERROR_INFO"

// ErrorHandler 最外层错误处理中间件
// 捕获所有错误并显示友好的错误页面
func ErrorHandler(r *ghttp.Request) {
	// 继续处理请求
	r.Middleware.Next()

	// 检查响应状态码，如果是错误状态码则显示错误页面
	status := r.Response.Status
	if status >= 400 && status < 600 {
		// 尝试从上下文中获取自定义错误信息
		if errorInfo := r.GetCtxVar(ctxKeyErrorInfo); !errorInfo.IsNil() {
			// Handler 已经调用了 RenderError，不再处理
			return
		}

		// 获取已写入的响应内容作为错误消息
		message := r.Response.BufferString()
		if message == "" {
			message = http.StatusText(status)
		}

		// 清空响应缓冲区
		r.Response.ClearBuffer()

		// 显示错误页面
		showErrorPage(r, status, message, "")
	}
}

// RenderError 主动渲染错误页面（供 Handler 调用）
// 使用场景：需要更详细的错误信息或自定义错误展示
func RenderError(r *ghttp.Request, info ErrorInfo) {
	// 标记已处理，避免中间件重复处理
	r.SetCtxVar(ctxKeyErrorInfo, true)

	// 如果没有设置状态码，默认使用 500
	if info.Code == 0 {
		info.Code = http.StatusInternalServerError
	}

	// 如果没有设置消息，使用标准 HTTP 状态文本
	if info.Message == "" {
		info.Message = http.StatusText(info.Code)
	}

	// 清空已有响应缓冲区
	r.Response.ClearBuffer()

	// 渲染错误页面
	showErrorPage(r, info.Code, info.Message, info.Detail, info.CustomJS)
}

// showErrorPage 显示错误页面
func showErrorPage(r *ghttp.Request, statusCode int, message string, detail string, customJS ...string) {
	ctx := r.Context()

	// 记录错误日志
	logError(ctx, statusCode, message, detail)

	// 设置响应头
	r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.Response.Status = statusCode

	// 获取自定义 JS（可选参数）
	js := ""
	if len(customJS) > 0 {
		js = customJS[0]
	}

	// 渲染错误页面
	htmlContent := generateErrorHTML(statusCode, message, detail, js)
	r.Response.Write(htmlContent)
}

// logError 记录错误日志
func logError(ctx context.Context, statusCode int, message string, detail string) {
	logger := g.Log()
	logMsg := fmt.Sprintf("错误码: %d, 错误信息: %s", statusCode, message)

	if detail != "" {
		logMsg += fmt.Sprintf(", 详细信息: %s", detail)
	}

	// 根据状态码选择日志级别
	switch {
	case statusCode >= 500:
		logger.Error(ctx, logMsg)
	case statusCode >= 400:
		logger.Warning(ctx, logMsg)
	default:
		logger.Info(ctx, logMsg)
	}
}

// generateErrorHTML 生成错误页面HTML
func generateErrorHTML(statusCode int, message string, detail string, customJS string) string {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>%d %s</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
        padding-top: 50px;
    }
    h1 {
        font-size: 1.5em;
        font-weight: normal;
        margin: 0 0 0.5em 0;
        padding: 0;
    }
    p {
        margin: 0.5em 0;
        line-height: 1.6;
    }
    a {
        color: #06c;
        text-decoration: none;
    }
    a:hover {
        text-decoration: underline;
    }
    .detail {
        margin-top: 1.5em;
        padding: 10px;
        background-color: #f0f0f0;
        border-left: 3px solid #999;
        font-size: 0.9em;
        color: #666;
    }
    hr {
        border: 0;
        border-top: 1px solid #ccc;
        margin: 1.5em 0;
    }
    .footer {
        margin-top: 2em;
        font-size: 0.9em;
        color: #999;
    }
</style>
%s
</head>
<body>
<h1>%d %s</h1>
<p>%s</p>
%s
<hr>
<p class="footer">香港中文大学（深圳）GPT服务统一鉴权系统</p>
</body>
</html>`, statusCode, http.StatusText(statusCode), generateCustomJSSection(customJS), statusCode, http.StatusText(statusCode), message, generateDetailSection(detail))

	return html
}

// generateCustomJSSection 生成自定义 JS 部分
func generateCustomJSSection(customJS string) string {
	if customJS == "" {
		return ""
	}
	return fmt.Sprintf("<script>\n%s\n</script>", customJS)
}

// generateDetailSection 生成详细信息部分
func generateDetailSection(detail string) string {
	if detail == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="detail">
<strong>详细信息：</strong><br>%s
</div>`, escapeHTML(detail))
}

// escapeHTML 简单的HTML转义
func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}
