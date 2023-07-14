package tools

import (
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/api"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

type ApiClient struct {
	api.BaseApiClient
	common service

	// API Services

	AdminApi   *ToolAdminApiService
	CountryApi *ToolCountryApiService
	RegionApi  *ToolRegionApiService
}

type service struct {
	client *ApiClient
}

// NewApiClient 创建新的Tools API客户端
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

	c.AdminApi = (*ToolAdminApiService)(&c.common)
	c.CountryApi = (*ToolCountryApiService)(&c.common)
	c.RegionApi = (*ToolRegionApiService)(&c.common)

	return c
}
