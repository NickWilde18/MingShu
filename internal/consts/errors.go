package consts

import "net/http"

// 错误代码定义
// 使用自定义错误代码，便于前后端对接和日志追踪
const (
	// 通用错误 (1xxx)
	ErrCodeUnknown            = 1000 // 未知错误
	ErrCodeInternalServer     = 1001 // 内部服务器错误
	ErrCodeServiceUnavailable = 1002 // 服务不可用
	ErrCodeTimeout            = 1003 // 请求超时

	// 认证相关错误 (2xxx)
	ErrCodeUnauthorized   = 2001 // 未认证
	ErrCodeSessionExpired = 2002 // 会话过期
	ErrCodeInvalidToken   = 2003 // 无效的令牌
	ErrCodeLoginFailed    = 2004 // 登录失败
	ErrCodeLogoutFailed   = 2005 // 登出失败
	ErrCodeCallbackFailed = 2006 // SSO回调失败

	// 权限相关错误 (3xxx)
	ErrCodeForbidden = 3001 // 无权访问

	// 请求相关错误 (4xxx)
	ErrCodeBadRequest            = 4001 // 请求参数错误
	ErrCodeNotFound              = 4004 // 资源未找到
	ErrCodeMethodNotAllowed      = 4005 // 方法不允许
	ErrCodeInvalidParameter      = 4006 // 参数无效
	ErrCodeTooManyRequests       = 4007 // 请求过多
	ErrCodeRequestEntityTooLarge = 4008 // 请求实体过大

	// 代理相关错误 (5xxx)
	ErrCodeBadGateway      = 5001 // 代理失败
	ErrCodeUpstreamTimeout = 5002 // 上游服务超时
)

// ErrorCode 错误码结构
type ErrorCode struct {
	Code          int    // 自定义错误代码
	HTTPStatus    int    // HTTP状态码
	TitleZh       string // 中文标题
	TitleEn       string // 英文标题
	MessageZh     string // 中文消息
	MessageEn     string // 英文消息
	SuggestionZh  string // 中文建议
	SuggestionEn  string // 英文建议
	ButtonLeft    string // 自定义左按钮文本
	ButtonLeftJS  string // 自定义左按钮JS
	ButtonRight   string // 自定义右按钮文本
	ButtonRightJS string // 自定义右按钮JS
}

// 根据 HTTP 状态码映射默认的错误码
var DefaultErrorCodeMap = map[int]int{
	http.StatusInternalServerError:   ErrCodeInternalServer,
	http.StatusServiceUnavailable:    ErrCodeServiceUnavailable,
	http.StatusRequestTimeout:        ErrCodeTimeout,
	http.StatusUnauthorized:          ErrCodeUnauthorized,
	http.StatusForbidden:             ErrCodeForbidden,
	http.StatusBadRequest:            ErrCodeBadRequest,
	http.StatusNotFound:              ErrCodeNotFound,
	http.StatusMethodNotAllowed:      ErrCodeMethodNotAllowed,
	http.StatusBadGateway:            ErrCodeBadGateway,
	http.StatusGatewayTimeout:        ErrCodeUpstreamTimeout,
	http.StatusRequestEntityTooLarge: ErrCodeRequestEntityTooLarge,
	http.StatusTooManyRequests:       ErrCodeTooManyRequests,
}

// 错误码映射表（中英文一起定义）
var ErrorCodeMap = map[int]ErrorCode{
	// 通用错误
	ErrCodeUnknown: {
		Code:          ErrCodeUnknown,
		HTTPStatus:    http.StatusInternalServerError,
		TitleZh:       "未知错误",
		TitleEn:       "Unknown Error",
		MessageZh:     "发生了未知错误，请稍后重试。",
		MessageEn:     "An unknown error occurred. Please try again later.",
		SuggestionZh:  "如果问题持续存在，请携带追踪ID联系ITSO。",
		SuggestionEn:  "If the problem persists, please contact technical support with the trace ID.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeInternalServer: {
		Code:          ErrCodeInternalServer,
		HTTPStatus:    http.StatusInternalServerError,
		TitleZh:       "内部服务器错误",
		TitleEn:       "Internal Server Error",
		MessageZh:     "服务器遇到了意外情况，我们正在努力修复。",
		MessageEn:     "The server encountered an unexpected condition. We are working to fix it.",
		SuggestionZh:  "请稍等几分钟后重试。如果问题持续，请联系ITSO。",
		SuggestionEn:  "Please try again in a few minutes. If the issue continues, contact support.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeServiceUnavailable: {
		Code:          ErrCodeServiceUnavailable,
		HTTPStatus:    http.StatusServiceUnavailable,
		TitleZh:       "服务不可用",
		TitleEn:       "Service Unavailable",
		MessageZh:     "服务因维护或高负载暂时不可用。",
		MessageEn:     "The service is temporarily unavailable due to maintenance or high load.",
		SuggestionZh:  "请稍候片刻后重试。您也可以查看我们的状态页面获取更新。",
		SuggestionEn:  "Please wait a moment and try again. Check our status page for updates.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeTimeout: {
		Code:          ErrCodeTimeout,
		HTTPStatus:    http.StatusRequestTimeout,
		TitleZh:       "请求超时",
		TitleEn:       "Request Timeout",
		MessageZh:     "请求处理时间过长，已超时。",
		MessageEn:     "The request took too long to process and timed out.",
		SuggestionZh:  "请检查您的网络连接并重试。",
		SuggestionEn:  "Please check your network connection and try again.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},

	// 认证相关
	ErrCodeUnauthorized: {
		Code:          ErrCodeUnauthorized,
		HTTPStatus:    http.StatusUnauthorized,
		TitleZh:       "需要身份验证",
		TitleEn:       "Authentication Required",
		MessageZh:     "您需要登录才能访问此资源。",
		MessageEn:     "You need to log in to access this resource.",
		SuggestionZh:  "请点击下方按钮进行登录。",
		SuggestionEn:  "Please click the button below to log in.",
		ButtonLeft:    "前往登录页 / Login Page",
		ButtonLeftJS:  "location.href = '/auth/login';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeSessionExpired: {
		Code:          ErrCodeSessionExpired,
		HTTPStatus:    http.StatusUnauthorized,
		TitleZh:       "会话已过期",
		TitleEn:       "Session Expired",
		MessageZh:     "出于安全考虑，您的会话已过期。",
		MessageEn:     "Your session has expired for security reasons.",
		SuggestionZh:  "请重新登录以继续操作。",
		SuggestionEn:  "Please log in again to continue.",
		ButtonLeft:    "前往登录页 / Login Page",
		ButtonLeftJS:  "location.href = '/auth/login';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeInvalidToken: {
		Code:          ErrCodeInvalidToken,
		HTTPStatus:    http.StatusUnauthorized,
		TitleZh:       "无效的令牌",
		TitleEn:       "Invalid Token",
		MessageZh:     "身份验证令牌无效或已被篡改。",
		MessageEn:     "The authentication token is invalid or has been tampered with.",
		SuggestionZh:  "请重新登录以获取新的令牌。",
		SuggestionEn:  "Please log in again to get a new token.",
		ButtonLeft:    "前往登录页 / Login Page",
		ButtonLeftJS:  "location.href = '/auth/login';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeLoginFailed: {
		Code:          ErrCodeLoginFailed,
		HTTPStatus:    http.StatusUnauthorized,
		TitleZh:       "登录失败",
		TitleEn:       "Login Failed",
		MessageZh:     "无法完成登录流程。",
		MessageEn:     "Unable to complete the login process.",
		SuggestionZh:  "请验证您的凭据并重试。",
		SuggestionEn:  "Please verify your credentials and try again.",
		ButtonLeft:    "前往登录页 / Login Page",
		ButtonLeftJS:  "location.href = '/auth/login';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeLogoutFailed: {
		Code:         ErrCodeLogoutFailed,
		HTTPStatus:   http.StatusInternalServerError,
		TitleZh:      "登出失败",
		TitleEn:      "Logout Failed",
		MessageZh:    "登出过程中发生错误。",
		MessageEn:    "An error occurred during the logout process.",
		SuggestionZh: "您可以直接关闭此窗口或重试。",
		SuggestionEn: "You may close this window directly or try again.",
		// 只留一个按钮
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeCallbackFailed: {
		Code:          ErrCodeCallbackFailed,
		HTTPStatus:    http.StatusInternalServerError,
		TitleZh:       "SSO回调失败",
		TitleEn:       "SSO Callback Failed",
		MessageZh:     "处理SSO身份验证回调时失败。",
		MessageEn:     "Failed to process the SSO authentication callback.",
		SuggestionZh:  "请从头开始重新登录。",
		SuggestionEn:  "Please try logging in again from the beginning.",
		ButtonLeft:    "前往登录页 / Login Page",
		ButtonLeftJS:  "location.href = '/auth/login';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},

	// 权限相关
	ErrCodeForbidden: {
		Code:          ErrCodeForbidden,
		HTTPStatus:    http.StatusForbidden,
		TitleZh:       "禁止访问",
		TitleEn:       "Access Forbidden",
		MessageZh:     "您没有权限访问此资源。",
		MessageEn:     "You don't have permission to access this resource.",
		SuggestionZh:  "如果您认为应该有访问权限，请联系ITSO。",
		SuggestionEn:  "If you believe you should have access, please contact your administrator.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},

	// 请求相关
	ErrCodeBadRequest: {
		Code:          ErrCodeBadRequest,
		HTTPStatus:    http.StatusBadRequest,
		TitleZh:       "错误的请求",
		TitleEn:       "Bad Request",
		MessageZh:     "由于语法或参数无效，请求无法处理。",
		MessageEn:     "The request cannot be processed due to invalid syntax or parameters.",
		SuggestionZh:  "请检查您的请求后重试。",
		SuggestionEn:  "Please check your request and try again.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeNotFound: {
		Code:          ErrCodeNotFound,
		HTTPStatus:    http.StatusNotFound,
		TitleZh:       "页面未找到",
		TitleEn:       "Page Not Found",
		MessageZh:     "您访问的页面不存在或已被移动。",
		MessageEn:     "The page you are looking for doesn't exist or has been moved.",
		SuggestionZh:  "请检查URL或返回首页。",
		SuggestionEn:  "Please check the URL or return to the home page.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeMethodNotAllowed: {
		Code:          ErrCodeMethodNotAllowed,
		HTTPStatus:    http.StatusMethodNotAllowed,
		TitleZh:       "方法不被允许",
		TitleEn:       "Method Not Allowed",
		MessageZh:     "此资源不支持使用的HTTP方法。",
		MessageEn:     "The HTTP method used is not supported for this resource.",
		SuggestionZh:  "请检查您的请求方法后重试。",
		SuggestionEn:  "Please check your request method and try again.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeInvalidParameter: {
		Code:          ErrCodeInvalidParameter,
		HTTPStatus:    http.StatusBadRequest,
		TitleZh:       "参数无效",
		TitleEn:       "Invalid Parameter",
		MessageZh:     "您的请求中包含一个或多个无效参数。",
		MessageEn:     "One or more parameters in your request are invalid.",
		SuggestionZh:  "请验证所有参数及其值。",
		SuggestionEn:  "Please verify all parameters and their values.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeTooManyRequests: {
		Code:          ErrCodeTooManyRequests,
		HTTPStatus:    http.StatusTooManyRequests,
		TitleZh:       "请求过多",
		TitleEn:       "Too Many Requests",
		MessageZh:     "您的请求过于频繁，请稍后再试。",
		MessageEn:     "Your request is too frequent. Please try again later.",
		SuggestionZh:  "请稍后再试。",
		SuggestionEn:  "Please try again later.",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},

	// 代理相关
	ErrCodeUpstreamTimeout: {
		Code:          ErrCodeUpstreamTimeout,
		HTTPStatus:    http.StatusGatewayTimeout,
		TitleZh:       "上游服务超时",
		TitleEn:       "Upstream Service Timeout",
		MessageZh:     "上游服务未能及时响应。",
		MessageEn:     "The upstream service did not respond in time.",
		SuggestionZh:  "服务可能正在经历高负载，请稍后重试。",
		SuggestionEn:  "The service may be experiencing high load. Please try again later.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeBadGateway: {
		Code:          ErrCodeBadGateway,
		HTTPStatus:    http.StatusBadGateway,
		TitleZh:       "网关错误",
		TitleEn:       "Bad Gateway",
		MessageZh:     "我们这边似乎出了一点技术问题，导致您的请求无法完成。",
		MessageEn:     "It seems we're experiencing a technical issue on our end that prevented your request from completing.",
		SuggestionZh:  "这通常是暂时的，请稍候重试。",
		SuggestionEn:  "This is usually temporary. Please try again in a moment.",
		ButtonLeft:    "返回首页 / Home Page",
		ButtonLeftJS:  "location.href = '/';",
		ButtonRight:   "重试 / Retry",
		ButtonRightJS: "location.reload();",
	},
	ErrCodeRequestEntityTooLarge: {
		Code:         ErrCodeRequestEntityTooLarge,
		HTTPStatus:   http.StatusRequestEntityTooLarge,
		TitleZh:      "请求实体过大",
		TitleEn:      "Request Entity Too Large",
		MessageZh:    "您的请求实体过大，请尝试缩小请求体大小。",
		MessageEn:    "Your request entity is too large. Please try to reduce the request body size.",
		SuggestionZh: "请缩小请求体大小后重试。",
		SuggestionEn: "Please reduce the request body size and try again.",
	},
}
