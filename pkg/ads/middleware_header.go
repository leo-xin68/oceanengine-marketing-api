package ads

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

// AuthMiddleware ...
type AuthMiddleware struct {
	config *SdkConfig
}

// Handle ...
func (a *AuthMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {

	cfg := a.config
	if cfg.GlobalConfig.HttpOption.Header != nil {
		for k, v := range cfg.GlobalConfig.HttpOption.Header {
			req.Header[k] = v
		}
	}

	// 设置 AccessToken
	if cfg.AccessToken != "" {
		req.Header.Add("Access-Token", cfg.AccessToken)
	}

	// 设置 Debugger模式
	if cfg.IsDebug {
		req.Header.Add("X-Debug-Mode", "1")
	}

	rsp, err = next(req)
	return rsp, err
}
