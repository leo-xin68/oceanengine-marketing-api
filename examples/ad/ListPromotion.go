package main

import (
	"encoding/json"
	"fmt"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/examples"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/ads"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/util"
	"net/http"
	"strings"
)

type PromotionGetExample struct {
	OceanAds            *ads.App
	PromotionGetRequest model.PromotionGetRequest
}

func (e *PromotionGetExample) Init() {
	cfg := config.NewSdkConfig(AppId, AccessToken, true)
	e.OceanAds = ads.Init(cfg)
}

func (e *PromotionGetExample) RunExample() (model.PromotionGetResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Promotion.List(ctx, e.PromotionGetRequest)
}

func main() {
	e := &PromotionGetExample{}
	e.Init()

	requestJsonDoc := strings.Replace(`
		{
			"advertiser_id":{advertiser_id},
			"filtering":{
				"promotion_create_time":"2023-07-15"
			},
			"page":1,
			"page_size":10
		}
	`, "{advertiser_id}", AdvertiserId, -1)

	json.Unmarshal([]byte(requestJsonDoc), &e.PromotionGetRequest)

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
