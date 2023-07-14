package ads

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/api/tools"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

// InitTools ...
func InitTools(cfg *config.SdkConfig, ctx *context.Context) *Tools {

	if ctx == nil {
		_ctx := context.Background()
		_ctx = context.WithValue(_ctx, "ContextAPIKey", cfg.AccessToken)
		ctx = &_ctx
	}

	client := tools.NewApiClient(cfg)
	t := &Tools{
		SdkClient: &SdkClient{
			Config:       cfg,
			Ctx:          ctx,
			RoundTripper: http.DefaultTransport,
			middlewareList: []Middleware{
				&AuthMiddleware{cfg},
			},
		},
	}

	client.Cfg.HttpClient.Transport = t

	t.Admin = client.AdminApi
	t.Country = client.CountryApi
	t.Region = client.RegionApi

	return t
}

type Tools struct {
	*SdkClient

	Admin   *tools.ToolAdminApiService
	Country *tools.ToolCountryApiService
	Region  *tools.ToolRegionApiService
}
