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
	Code    int
	Message string
	Detail  string // 详细信息（可选）
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
	showErrorPage(r, info.Code, info.Message, info.Detail)
}

// showErrorPage 显示错误页面
func showErrorPage(r *ghttp.Request, statusCode int, message string, detail string) {
	ctx := r.Context()

	// 设置响应状态码
	r.Response.WriteStatus(statusCode)

	// 记录错误日志
	logError(ctx, statusCode, message, detail)

	// 渲染错误页面
	htmlContent := generateErrorHTML(statusCode, message, detail)
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
func generateErrorHTML(statusCode int, message string, detail string) string {
	// 简洁美观的错误页面
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>错误 %d</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .error-container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
            max-width: 600px;
            width: 100%%;
            padding: 40px;
            animation: slideIn 0.5s ease-out;
        }
        @keyframes slideIn {
            from {
                opacity: 0;
                transform: translateY(-20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        .error-icon {
            text-align: center;
            font-size: 80px;
            margin-bottom: 20px;
        }
        .error-code {
            text-align: center;
            font-size: 48px;
            font-weight: bold;
            color: #667eea;
            margin-bottom: 10px;
        }
        .error-message {
            text-align: center;
            font-size: 20px;
            color: #333;
            margin-bottom: 30px;
            font-weight: 500;
        }
        .detail-section {
            background: #fff3cd;
            border-left: 4px solid #ffc107;
            padding: 15px;
            margin-top: 20px;
            border-radius: 5px;
            font-size: 14px;
            color: #856404;
            word-break: break-all;
        }
        .back-button {
            display: block;
            width: 100%%;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            text-align: center;
            border-radius: 10px;
            text-decoration: none;
            font-weight: 600;
            font-size: 16px;
            margin-top: 20px;
            transition: transform 0.2s, box-shadow 0.2s;
        }
        .back-button:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 20px rgba(102, 126, 234, 0.3);
        }
    </style>
</head>
<body>
    <div class="error-container">
        <div class="error-icon">⚠️</div>
        <div class="error-code">错误 %d</div>
        <div class="error-message">%s</div>
        %s
        <a href="javascript:history.back()" class="back-button">返回上一页</a>
    </div>
</body>
</html>`, statusCode, statusCode, escapeHTML(message), generateDetailSection(detail))

	return html
}

// generateDetailSection 生成详细信息部分
func generateDetailSection(detail string) string {
	if detail == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="detail-section">
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
