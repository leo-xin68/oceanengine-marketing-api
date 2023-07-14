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

type ToolAdminInfoGetExample struct {
	OceanAds             *ads.App
	ToolAdminGetInfoOpts *model.ToolAdminGetInfoOpts
}

func (e *ToolAdminInfoGetExample) Init() {
	cfg := config.NewSdkConfig(AppId, AccessToken, true)
	e.OceanAds = ads.Init(cfg)
}

func (e *ToolAdminInfoGetExample) RunExample() (model.ToolAdminInfoGetResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Tools.Admin.Info(ctx, e.ToolAdminGetInfoOpts)
}

func main() {
	e := &ToolAdminInfoGetExample{}
	e.Init()

	e.ToolAdminGetInfoOpts = &model.ToolAdminGetInfoOpts{
		AdvertiserId: optional.NewString(AdvertiserId),
		Codes:        optional.NewInterface([]string{"CN"}),
		Language:     optional.NewString("ZH_CN"),
		SubDistrict:  optional.NewString("FOUR_LEVEL"),
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
