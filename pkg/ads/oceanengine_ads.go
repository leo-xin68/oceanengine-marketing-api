package ads

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/api"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

// SdkClient ...
type SdkClient struct {
	http.RoundTripper
	Config         *config.SdkConfig
	Client         *api.ApiClient
	Ctx            *context.Context
	Version        string
	middlewareList []Middleware
}

// RoundTrip ...
func (c *SdkClient) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(req *http.Request) (rsp *http.Response, err error) {
		return c.RoundTripper.RoundTrip(req)
	}
	middlewareCount := len(c.middlewareList)
	// 逆序遍历
	for i := middlewareCount - 1; i >= 0; i-- {
		beforeFunc = c.GenMiddlewareHandleFunc(c.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}

// GenMiddlewareHandleFunc ...
func (c *SdkClient) GenMiddlewareHandleFunc(
	middleware Middleware,
	beforeFunc func(req *http.Request) (rsp *http.Response, err error),
) func(req *http.Request) (rsp *http.Response, err error) {
	return func(req *http.Request) (rsp *http.Response, err error) {
		return middleware.Handle(req, beforeFunc)
	}
}

// Init ...
func Init(cfg *config.SdkConfig) *SdkClient {
	cfg.BasePath = "https://ad.oceanengine.com/open_api"
	version := "1.0.0"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ContextAPIKey", cfg.AccessToken)
	client := api.NewApiClient(cfg)
	sdkClient := &SdkClient{
		Config:       cfg,
		Ctx:          &ctx,
		Client:       client,
		RoundTripper: http.DefaultTransport,
		Version:      version,
	}

	sdkClient.Client.Cfg.HttpClient.Transport = sdkClient
	sdkClient.middlewareList = []Middleware{
		&AuthMiddleware{sdkClient},
	}
	return sdkClient
}
