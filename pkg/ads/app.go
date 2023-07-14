package ads

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/api"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"net/http"
)

// Init ...
func Init(cfg *SdkConfig) *App {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ContextAPIKey", cfg.AccessToken)
	client := api.NewApiClient(cfg)

	ads := &App{
		SdkClient: &SdkClient{
			Config:       cfg,
			Ctx:          &ctx,
			RoundTripper: http.DefaultTransport,
		},
	}
	ads.SdkClient.middlewareList = []Middleware{
		&AuthMiddleware{cfg},
	}
	client.Cfg.HttpClient.Transport = ads

	// init tools
	ads.Tools = InitTools(cfg, &ctx)

	ads.Project = client.ProjectApi

	return ads
}

type App struct {
	*SdkClient
	Tools *Tools

	Project *api.ProjectApiService
}
