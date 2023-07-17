package api

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

// ApiClient 负责请求api。在大多数情况下，应该只有一个共享的APIClient
type ApiClient struct {
	BaseApiClient
	common service // 复用单个结构，而不是为每个服务分配一个结构

	// API Services
	ProjectApi   *ProjectApiService
	PromotionApi *PromotionApiService
}

type service struct {
	client *ApiClient
}

// NewApiClient 创建新的API客户端。需要描述应用程序的userAgent字符串。
// 可选的自定义http.Client，来支持缓存之类的高级功能。
func NewApiClient(sdkConfig *SdkConfig) *ApiClient {
	cfg := sdkConfig.Configuration
	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{}
	}

	c := &ApiClient{}
	c.SdkConfig = sdkConfig
	c.Cfg = &cfg
	c.common.client = c

	// API Services
	c.ProjectApi = (*ProjectApiService)(&c.common)
	c.PromotionApi = (*PromotionApiService)(&c.common)

	return c
}
