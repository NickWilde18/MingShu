package service

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

var proxyHostMap = map[string]string{
	"www.baidu.com": "http://www.baidu.com/",
	"uniauth-gf": "http://localhost:8000/",
	"uniauth-admin": "http://localhost:5173/",
}

func GetProxyHost(service string) (string, error) {
	proxyHost, ok := proxyHostMap[service]
	if !ok {
		return "", gerror.Newf("未找到 %s 服务。请确认请求头是否传递正确的服务名称！", service)
	}
	return proxyHost, nil
}