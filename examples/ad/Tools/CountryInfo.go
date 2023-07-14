package main

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/examples"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/ads"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	model "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/tools"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/util"
	"net/http"
)

type ToolCountryInfoGetExample struct {
	OceanAds               *ads.App
	ToolCountryGetInfoOpts *model.ToolCountryGetInfoOpts
}

func (e *ToolCountryInfoGetExample) Init() {
	cfg := config.NewSdkConfig(AppId, AccessToken, true)
	e.OceanAds = ads.Init(cfg)
}

func (e *ToolCountryInfoGetExample) RunExample() (model.ToolCountryInfoGetResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Tools.Country.Info(ctx, e.ToolCountryGetInfoOpts)
}

func main() {
	e := &ToolCountryInfoGetExample{}
	e.Init()
	e.ToolCountryGetInfoOpts = &model.ToolCountryGetInfoOpts{
		AdvertiserId: optional.NewString(AdvertiserId),
		Language:     optional.NewString("ZH_CN"),
	}

	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}

	util.JsonFormatPrint("Response data", response)
	util.JsonPrint("Headers", headers)
}
