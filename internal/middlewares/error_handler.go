package middlewares

import (
	_ "embed"
	"net/http"
	"time"

	"uniauth-gateway/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

//go:embed error.tpl
var errorTemplate string

// ErrorInfo 错误信息结构
type ErrorInfo struct {
	// 错误代码，主要由这个决定显示的内容，必须传。
	ErrorCode int
	// 对应的错误说明中英文（MessageZH/EN）不允许动态传递。如果要加必须改代码，新增错误码。
	// 对应的建议中英文（SuggestionZH/EN）不允许动态传递。如果要加必须改代码，新增错误码。

	// 报错详情（可选，默认显示。可以在配置文件中开关 showDetail）
	Detail string
	// 自定义消息（可选，需要中英文都写好）
	CustomMsg string
	// 自定义JavaScript代码（可选），会注入到错误页面的 <script> 标签中
	CustomJS string
}

const ctxKeyErrorInfo = "ERROR_INFO"

// ErrorHandler 自定义错误处理
// 注意：该函数执行后会在**自动**退出当前 Handler。
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 错误兜底
	status := r.Response.Status
	if status >= 400 && status < 600 {
		if errorInfo := r.GetCtxVar(ctxKeyErrorInfo); !errorInfo.IsNil() {
			return
		}
		errorCode, exists := consts.DefaultErrorCodeMap[status]
		if !exists {
			errorCode = consts.ErrCodeUnknown
		}
		RenderError(r, ErrorInfo{
			ErrorCode: errorCode,
			Detail:    r.Response.BufferString(),
		})
	}
}

// RenderError 渲染错误页面
func RenderError(r *ghttp.Request, info ErrorInfo) {
	ctx := r.GetCtx()
	r.SetCtxVar(ctxKeyErrorInfo, true)
	r.Response.ClearBuffer()
	r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 错误模版要求提供：
	// 1. HTTPStatus
	// 2. TitleZh
	// 3. TitleEn
	// 4. MessageZh
	// 5. MessageEn
	// 6. CustomMsg
	// 7. SuggestionZh
	// 8. SuggestionEn
	// 9. Detail
	// 10. ShowDetail
	// 11. TraceID
	// 12. TraceIDShort
	// 13. ErrCode
	// 14. Timestamp
	// 15. ButtonLeft
	// 16. ButtonLeftJS
	// 17. ButtonRight
	// 18. ButtonRightJS
	// 19. CustomJS

	// 获取错误码配置
	errorCodeConfig, exists := consts.ErrorCodeMap[info.ErrorCode]
	if !exists {
		errorCodeConfig = consts.ErrorCodeMap[consts.ErrCodeUnknown]
		info.ErrorCode = consts.ErrCodeUnknown
	}

	// 构建模板数据
	data := g.Map{
		"HTTPStatus":   errorCodeConfig.HTTPStatus,
		"TitleZh":      errorCodeConfig.TitleZh,
		"TitleEn":      errorCodeConfig.TitleEn,
		"MessageZh":    errorCodeConfig.MessageZh,
		"MessageEn":    errorCodeConfig.MessageEn,
		"SuggestionZh": errorCodeConfig.SuggestionZh,
		"SuggestionEn": errorCodeConfig.SuggestionEn,

		"ShowDetail":   g.Cfg().MustGet(ctx, "server.showDetail", true).Bool(),
		"Detail":       escapeHTML(info.Detail),
		"TraceID":      gctx.CtxId(ctx),
		"TraceIDShort": gctx.CtxId(ctx)[:7] + "...",
		"ErrorCode":    errorCodeConfig.Code,
		"Timestamp":    time.Now().Format("2006-01-02 15:04:05"),

		"CustomMsg":     info.CustomMsg,
		"CustomJS":      info.CustomJS,
		"ButtonLeft":    errorCodeConfig.ButtonLeft,
		"ButtonLeftJS":  errorCodeConfig.ButtonLeftJS,
		"ButtonRight":   errorCodeConfig.ButtonRight,
		"ButtonRightJS": errorCodeConfig.ButtonRightJS,
	}

	// 使用 GoFrame 模板引擎渲染
	r.Response.Status = errorCodeConfig.HTTPStatus
	err := r.Response.WriteTplContent(errorTemplate, data)
	if err != nil {
		// 如果模板渲染失败，返回一个简单的错误页面
		g.Log().Error(ctx, "错误页面模板渲染失败:", err)
		r.Response.ClearBuffer()
		r.Response.Writeln("Trace ID:", gctx.CtxId(ctx))
		r.Response.Writeln(gerror.Wrap(err, "错误页面模板渲染失败"))
		r.Response.Status = http.StatusBadGateway
	}
	r.Exit()
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
