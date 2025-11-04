package consts

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
	ErrCodeForbidden    = 3001 // 无权访问
	ErrCodeNoPermission = 3002 // 没有操作权限
	ErrCodeAccessDenied = 3003 // 访问被拒绝

	// 请求相关错误 (4xxx)
	ErrCodeBadRequest       = 4001 // 请求参数错误
	ErrCodeNotFound         = 4004 // 资源未找到
	ErrCodeMethodNotAllowed = 4005 // 方法不允许
	ErrCodeInvalidParameter = 4006 // 参数无效

	// 代理相关错误 (5xxx)
	ErrCodeProxyFailed     = 5001 // 代理失败
	ErrCodeUpstreamTimeout = 5002 // 上游服务超时
	ErrCodeUpstreamError   = 5003 // 上游服务错误
	ErrCodeBadGateway      = 5004 // 网关错误
)

// ErrorCode 错误码结构
//
type ErrorCode struct {
	Code         int    // 自定义错误代码
	HTTPStatus   int    // HTTP状态码
	TitleZh      string // 中文标题
	TitleEn      string // 英文标题
	MessageZh    string // 中文消息
	MessageEn    string // 英文消息
	SuggestionZh string // 中文建议
	SuggestionEn string // 英文建议
}

// 错误码映射表（中英文一起定义）
var ErrorCodeMap = map[int]ErrorCode{
	// 通用错误
	ErrCodeUnknown: {
		Code:         ErrCodeUnknown,
		HTTPStatus:   500,
		TitleZh:      "未知错误",
		TitleEn:      "Unknown Error",
		MessageZh:    "发生了未知错误，请稍后重试。",
		MessageEn:    "An unknown error occurred. Please try again later.",
		SuggestionZh: "如果问题持续存在，请携带追踪ID联系技术支持。",
		SuggestionEn: "If the problem persists, please contact technical support with the trace ID.",
	},
	ErrCodeInternalServer: {
		Code:         ErrCodeInternalServer,
		HTTPStatus:   500,
		TitleZh:      "内部服务器错误",
		TitleEn:      "Internal Server Error",
		MessageZh:    "服务器遇到了意外情况，我们正在努力修复。",
		MessageEn:    "The server encountered an unexpected condition. We are working to fix it.",
		SuggestionZh: "请稍等几分钟后重试。如果问题持续，请联系支持团队。",
		SuggestionEn: "Please try again in a few minutes. If the issue continues, contact support.",
	},
	ErrCodeServiceUnavailable: {
		Code:         ErrCodeServiceUnavailable,
		HTTPStatus:   503,
		TitleZh:      "服务不可用",
		TitleEn:      "Service Unavailable",
		MessageZh:    "服务因维护或高负载暂时不可用。",
		MessageEn:    "The service is temporarily unavailable due to maintenance or high load.",
		SuggestionZh: "请稍候片刻后重试。您也可以查看我们的状态页面获取更新。",
		SuggestionEn: "Please wait a moment and try again. Check our status page for updates.",
	},
	ErrCodeTimeout: {
		Code:         ErrCodeTimeout,
		HTTPStatus:   504,
		TitleZh:      "请求超时",
		TitleEn:      "Request Timeout",
		MessageZh:    "请求处理时间过长，已超时。",
		MessageEn:    "The request took too long to process and timed out.",
		SuggestionZh: "请检查您的网络连接并重试。",
		SuggestionEn: "Please check your network connection and try again.",
	},

	// 认证相关
	ErrCodeUnauthorized: {
		Code:         ErrCodeUnauthorized,
		HTTPStatus:   401,
		TitleZh:      "需要身份验证",
		TitleEn:      "Authentication Required",
		MessageZh:    "您需要登录才能访问此资源。",
		MessageEn:    "You need to log in to access this resource.",
		SuggestionZh: "请点击下方按钮进行登录。",
		SuggestionEn: "Please click the button below to log in.",
	},
	ErrCodeSessionExpired: {
		Code:         ErrCodeSessionExpired,
		HTTPStatus:   401,
		TitleZh:      "会话已过期",
		TitleEn:      "Session Expired",
		MessageZh:    "出于安全考虑，您的会话已过期。",
		MessageEn:    "Your session has expired for security reasons.",
		SuggestionZh: "请重新登录以继续操作。",
		SuggestionEn: "Please log in again to continue.",
	},
	ErrCodeInvalidToken: {
		Code:         ErrCodeInvalidToken,
		HTTPStatus:   401,
		TitleZh:      "无效的令牌",
		TitleEn:      "Invalid Token",
		MessageZh:    "身份验证令牌无效或已被篡改。",
		MessageEn:    "The authentication token is invalid or has been tampered with.",
		SuggestionZh: "请重新登录以获取新的令牌。",
		SuggestionEn: "Please log in again to get a new token.",
	},
	ErrCodeLoginFailed: {
		Code:         ErrCodeLoginFailed,
		HTTPStatus:   401,
		TitleZh:      "登录失败",
		TitleEn:      "Login Failed",
		MessageZh:    "无法完成登录流程。",
		MessageEn:    "Unable to complete the login process.",
		SuggestionZh: "请验证您的凭据并重试。",
		SuggestionEn: "Please verify your credentials and try again.",
	},
	ErrCodeLogoutFailed: {
		Code:         ErrCodeLogoutFailed,
		HTTPStatus:   500,
		TitleZh:      "登出失败",
		TitleEn:      "Logout Failed",
		MessageZh:    "登出过程中发生错误。",
		MessageEn:    "An error occurred during the logout process.",
		SuggestionZh: "您可以直接关闭此窗口或重试。",
		SuggestionEn: "You may close this window directly or try again.",
	},
	ErrCodeCallbackFailed: {
		Code:         ErrCodeCallbackFailed,
		HTTPStatus:   500,
		TitleZh:      "SSO回调失败",
		TitleEn:      "SSO Callback Failed",
		MessageZh:    "处理SSO身份验证回调时失败。",
		MessageEn:    "Failed to process the SSO authentication callback.",
		SuggestionZh: "请从头开始重新登录。",
		SuggestionEn: "Please try logging in again from the beginning.",
	},

	// 权限相关
	ErrCodeForbidden: {
		Code:         ErrCodeForbidden,
		HTTPStatus:   403,
		TitleZh:      "禁止访问",
		TitleEn:      "Access Forbidden",
		MessageZh:    "您没有权限访问此资源。",
		MessageEn:    "You don't have permission to access this resource.",
		SuggestionZh: "如果您认为应该有访问权限，请联系管理员。",
		SuggestionEn: "If you believe you should have access, please contact your administrator.",
	},
	ErrCodeNoPermission: {
		Code:         ErrCodeNoPermission,
		HTTPStatus:   403,
		TitleZh:      "权限不足",
		TitleEn:      "Insufficient Permissions",
		MessageZh:    "您缺少执行此操作所需的权限。",
		MessageEn:    "You lack the necessary permissions to perform this action.",
		SuggestionZh: "请向管理员申请所需权限。",
		SuggestionEn: "Please request the required permissions from your administrator.",
	},
	ErrCodeAccessDenied: {
		Code:         ErrCodeAccessDenied,
		HTTPStatus:   403,
		TitleZh:      "访问被拒绝",
		TitleEn:      "Access Denied",
		MessageZh:    "您对此资源的访问已被拒绝。",
		MessageEn:    "Your access to this resource has been denied.",
		SuggestionZh: "如需帮助，请联系支持团队。",
		SuggestionEn: "Contact support if you need assistance.",
	},

	// 请求相关
	ErrCodeBadRequest: {
		Code:         ErrCodeBadRequest,
		HTTPStatus:   400,
		TitleZh:      "错误的请求",
		TitleEn:      "Bad Request",
		MessageZh:    "由于语法或参数无效，请求无法处理。",
		MessageEn:    "The request cannot be processed due to invalid syntax or parameters.",
		SuggestionZh: "请检查您的请求后重试。",
		SuggestionEn: "Please check your request and try again.",
	},
	ErrCodeNotFound: {
		Code:         ErrCodeNotFound,
		HTTPStatus:   404,
		TitleZh:      "页面未找到",
		TitleEn:      "Page Not Found",
		MessageZh:    "您访问的页面不存在或已被移动。",
		MessageEn:    "The page you are looking for doesn't exist or has been moved.",
		SuggestionZh: "请检查URL或返回首页。",
		SuggestionEn: "Please check the URL or return to the home page.",
	},
	ErrCodeMethodNotAllowed: {
		Code:         ErrCodeMethodNotAllowed,
		HTTPStatus:   405,
		TitleZh:      "方法不被允许",
		TitleEn:      "Method Not Allowed",
		MessageZh:    "此资源不支持使用的HTTP方法。",
		MessageEn:    "The HTTP method used is not supported for this resource.",
		SuggestionZh: "请检查您的请求方法后重试。",
		SuggestionEn: "Please check your request method and try again.",
	},
	ErrCodeInvalidParameter: {
		Code:         ErrCodeInvalidParameter,
		HTTPStatus:   400,
		TitleZh:      "参数无效",
		TitleEn:      "Invalid Parameter",
		MessageZh:    "您的请求中包含一个或多个无效参数。",
		MessageEn:    "One or more parameters in your request are invalid.",
		SuggestionZh: "请验证所有参数及其值。",
		SuggestionEn: "Please verify all parameters and their values.",
	},

	// 代理相关
	ErrCodeProxyFailed: {
		Code:         ErrCodeProxyFailed,
		HTTPStatus:   502,
		TitleZh:      "代理错误",
		TitleEn:      "Proxy Error",
		MessageZh:    "无法将您的请求代理到上游服务。",
		MessageEn:    "Failed to proxy your request to the upstream service.",
		SuggestionZh: "目标服务可能暂时不可用，请重试。",
		SuggestionEn: "The target service might be temporarily unavailable. Please try again.",
	},
	ErrCodeUpstreamTimeout: {
		Code:         ErrCodeUpstreamTimeout,
		HTTPStatus:   504,
		TitleZh:      "上游服务超时",
		TitleEn:      "Upstream Service Timeout",
		MessageZh:    "上游服务未能及时响应。",
		MessageEn:    "The upstream service did not respond in time.",
		SuggestionZh: "服务可能正在经历高负载，请稍后重试。",
		SuggestionEn: "The service may be experiencing high load. Please try again later.",
	},
	ErrCodeUpstreamError: {
		Code:         ErrCodeUpstreamError,
		HTTPStatus:   502,
		TitleZh:      "上游服务错误",
		TitleEn:      "Upstream Service Error",
		MessageZh:    "上游服务返回了错误。",
		MessageEn:    "The upstream service returned an error.",
		SuggestionZh: "请重试，如果问题持续存在，请联系支持。",
		SuggestionEn: "Please try again or contact support if the issue persists.",
	},
	ErrCodeBadGateway: {
		Code:         ErrCodeBadGateway,
		HTTPStatus:   502,
		TitleZh:      "网关错误",
		TitleEn:      "Bad Gateway",
		MessageZh:    "网关从上游服务器收到了无效响应。",
		MessageEn:    "The gateway received an invalid response from the upstream server.",
		SuggestionZh: "这通常是暂时的，请稍候重试。",
		SuggestionEn: "This is usually temporary. Please try again in a moment.",
	},
}
