package middlewares

import (
	_ "embed"
	"time"

	"uniauth-gateway/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

//go:embed error.tpl
var errorTemplate string

// ErrorInfo 错误信息结构
type ErrorInfo struct {
	ErrorCode  int // 自定义错误代码，主要由这个决定显示的内容
	HTTPStatus int // HTTP状态码（可选，默认从ErrorCode映射）

	CustomMsg string // 自定义消息（可选，会同时显示中英文）

	ShowDetail bool   // 是否显示详细信息
	Detail     string // 技术详情（可选，仅调试模式显示）

	CustomJS string // 自定义JavaScript代码（可选），会注入到错误页面的 <script> 标签中
}

const ctxKeyErrorInfo = "ERROR_INFO"

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	status := r.Response.Status
	if status >= 400 && status < 600 {
		if errorInfo := r.GetCtxVar(ctxKeyErrorInfo); !errorInfo.IsNil() {
			return
		}
		// 没有处理一般就是兜底了
		message := r.Response.BufferString()
		r.Response.ClearBuffer()

		// 兜底就按默认的错误码映射
		errorCode := mapHTTPStatusToErrorCode(status)

		RenderError(r, ErrorInfo{
			ErrorCode:  errorCode,
			HTTPStatus: status,
			Detail:     message,
		})
	}
}

// RenderError 渲染错误页面
func RenderError(r *ghttp.Request, info ErrorInfo) {
	ctx := r.GetCtx()
	r.SetCtxVar(ctxKeyErrorInfo, true)

	// 获取错误码配置
	errorCodeConfig, exists := consts.ErrorCodeMap[info.ErrorCode]
	if !exists {
		errorCodeConfig = consts.ErrorCodeMap[consts.ErrCodeUnknown]
		info.ErrorCode = consts.ErrCodeUnknown
	}

	// 设置HTTP状态码
	if info.HTTPStatus == 0 {
		info.HTTPStatus = errorCodeConfig.HTTPStatus
	}
	r.Response.Status = info.HTTPStatus

	// 检查是否为调试模式
	debugMode := g.Cfg().MustGet(ctx, "server.debug", false).Bool()
	showDetail := info.ShowDetail || debugMode

	// 构建模板数据
	data := g.Map{
		"ErrorCode":    info.ErrorCode,
		"HTTPStatus":   info.HTTPStatus,
		"TitleZh":      errorCodeConfig.TitleZh,
		"TitleEn":      errorCodeConfig.TitleEn,
		"MessageZh":    errorCodeConfig.MessageZh,
		"MessageEn":    errorCodeConfig.MessageEn,
		"SuggestionZh": errorCodeConfig.SuggestionZh,
		"SuggestionEn": errorCodeConfig.SuggestionEn,
		"CustomMsg":    info.CustomMsg,
		"Detail":       info.Detail,
		"ShowDetail":   showDetail,
		"TraceID":      gctx.CtxId(ctx),
		"Timestamp":    time.Now().Format("2006-01-02 15:04:05"),
		"CustomJS":     info.CustomJS,
	}

	// 渲染HTML
	r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 处理错误原因
	reasonText := ""
	if data["Detail"].(string) != "" && !data["ShowDetail"].(bool) {
		reasonText = data["Detail"].(string)
	}
	data["ReasonText"] = escapeHTML(reasonText)

	// TraceID - 只显示前16位，悬停显示完整ID 32位
	traceID := data["TraceID"].(string)
	traceIDShort := traceID
	if len(traceID) > 16 {
		traceIDShort = traceID[:16] + "..."
	}
	data["TraceIDShort"] = traceIDShort

	// 使用 GoFrame 模板引擎渲染
	err := r.Response.WriteTplContent(errorTemplate, data)
	if err != nil {
		// 如果模板渲染失败，返回一个简单的错误页面
		g.Log().Error(ctx, "Error rendering template:", err)
        r.Response.ClearBuffer()
		r.Response.Write("错误页面模板渲染错误，错误信息：" + err.Error())
	}
}

// escapeHTML 转义HTML特殊字符，防止XSS攻击
func escapeHTML(s string) string {
	s = gstr.Replace(s, "&", "&amp;")
	s = gstr.Replace(s, "<", "&lt;")
	s = gstr.Replace(s, ">", "&gt;")
	s = gstr.Replace(s, "\"", "&quot;")
	s = gstr.Replace(s, "'", "&#39;")
	return s
}

// mapHTTPStatusToErrorCode 将HTTP状态码映射到自定义错误代码
func mapHTTPStatusToErrorCode(httpStatus int) int {
	switch httpStatus {
	case 400:
		return consts.ErrCodeBadRequest
	case 401:
		return consts.ErrCodeUnauthorized
	case 403:
		return consts.ErrCodeForbidden
	case 404:
		return consts.ErrCodeNotFound
	case 405:
		return consts.ErrCodeMethodNotAllowed
	case 500:
		return consts.ErrCodeInternalServer
	case 502:
		return consts.ErrCodeBadGateway
	case 503:
		return consts.ErrCodeServiceUnavailable
	case 504:
		return consts.ErrCodeTimeout
	default:
		return consts.ErrCodeUnknown
	}
}
