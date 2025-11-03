# 错误处理使用示例

## 混合方案说明

我们提供了**两种方式**来处理错误，可以根据场景选择：

### 方式 1：自动处理（推荐用于大部分场景）

适用于**简单错误**，只需设置状态码和简单消息。中间件会自动拦截并显示错误页面。

```go
func SimpleHandler(r *ghttp.Request) {
    // 场景 1: 只设置状态码（使用标准 HTTP 错误文本）
    r.Response.WriteStatus(http.StatusNotFound)
    
    // 场景 2: 设置状态码 + 自定义消息
    r.Response.WriteStatus(http.StatusUnauthorized)
    r.Response.Write("用户未登录")
    
    // 场景 3: 使用 WriteStatusExit 立即返回
    r.Response.WriteStatusExit(http.StatusForbidden, "权限不足")
}
```

**优点：** 代码简洁，不会遗漏，统一风格

---

### 方式 2：主动调用（推荐用于复杂场景）

适用于需要**详细错误信息**或**特殊展示**的场景。

```go
import "MingShu/internal/middlewares"

func ComplexHandler(r *ghttp.Request) {
    // 检查用户权限
    if !hasPermission(r) {
        middlewares.RenderError(r, middlewares.ErrorInfo{
            Code:    http.StatusForbidden,
            Message: "您没有权限访问此资源",
            Detail:  "需要 admin 权限，当前用户权限：user",
        })
        return
    }
    
    // 业务逻辑...
}

func DetailedErrorHandler(r *ghttp.Request) {
    err := doSomething()
    if err != nil {
        middlewares.RenderError(r, middlewares.ErrorInfo{
            Code:    http.StatusInternalServerError,
            Message: "操作失败",
            Detail:  fmt.Sprintf("错误详情: %v", err),
        })
        return
    }
}
```

**优点：** 灵活性高，可以展示详细信息

---

## 实际场景示例

### 场景 1：登录验证中间件（简单）

```go
func VerifyLoginStatus(r *ghttp.Request) {
    userID, err := r.Session.Get("user_id")
    if err != nil {
        r.Response.WriteStatusExit(http.StatusInternalServerError, "Session获取错误")
        return
    }
    if userID.IsNil() {
        r.Response.WriteStatusExit(http.StatusUnauthorized, "未登录")
        return
    }
    r.Middleware.Next()
}
```

### 场景 2：文件上传 Handler（详细）

```go
func UploadHandler(r *ghttp.Request) {
    file := r.GetUploadFile("file")
    if file == nil {
        middlewares.RenderError(r, middlewares.ErrorInfo{
            Code:    http.StatusBadRequest,
            Message: "未找到上传的文件",
            Detail:  "请确保表单字段名为 'file'，文件大小不超过 10MB",
        })
        return
    }
    
    // 保存文件...
}
```

### 场景 3：API 代理（详细错误信息）

```go
func ProxyHandler(r *ghttp.Request) {
    resp, err := proxyRequest(r)
    if err != nil {
        middlewares.RenderError(r, middlewares.ErrorInfo{
            Code:    http.StatusBadGateway,
            Message: "上游服务请求失败",
            Detail:  fmt.Sprintf("目标服务: %s\n错误信息: %v", targetURL, err),
        })
        return
    }
    
    // 处理响应...
}
```

---

## 如何选择？

| 场景 | 推荐方式 | 原因 |
|-----|---------|-----|
| 简单的权限检查 | 自动处理 | 代码简洁 |
| 资源不存在 | 自动处理 | 标准错误 |
| 需要显示错误详情 | 主动调用 | 信息丰富 |
| 调试信息（开发环境）| 主动调用 | 可以显示堆栈 |
| 业务逻辑错误 | 主动调用 | 更好的用户体验 |
| 中间件中的错误 | 自动处理 | 统一风格 |

---

## 两种方式的协作

两种方式可以**无缝协作**：

- 中间件会检测 Handler 是否已经调用了 `RenderError`
- 如果已调用，中间件不会重复处理
- 如果未调用但有错误状态码，中间件自动兜底

**最佳实践：** 
- 默认使用自动处理（简单高效）
- 需要详细信息时使用主动调用（灵活定制）

