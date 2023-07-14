package ads

import (
	"net/http"
)

// AuthMiddleware ...
type AuthMiddleware struct {
	oceanAds *SdkClient
}

// Handle ...
func (a *AuthMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {

	oceanAds := a.oceanAds
	if oceanAds.Config.GlobalConfig.HttpOption.Header != nil {
		for k, v := range oceanAds.Config.GlobalConfig.HttpOption.Header {
			req.Header[k] = v
		}
	}

	// 设置 AccessToken
	if oceanAds.Config.AccessToken != "" {
		req.Header.Add("Access-Token", oceanAds.Config.AccessToken)
	}

	// 设置 Debugger模式
	if oceanAds.Config.IsDebug {
		req.Header.Add("X-Debug-Mode", "1")
	}

	rsp, err = next(req)
	return rsp, err
}
