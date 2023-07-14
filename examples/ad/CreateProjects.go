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

type ProjectsCreateExample struct {
	OceanAds             *ads.App
	ProjectCreateRequest model.ProjectCreateRequest
}

func (e *ProjectsCreateExample) Init() {
	cfg := config.NewSdkConfig(AppId, AccessToken, true)
	e.OceanAds = ads.Init(cfg)
}

func (e *ProjectsCreateExample) RunExample() (model.ProjectCreateResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Project.Create(ctx, e.ProjectCreateRequest)
}

func main() {
	e := &ProjectsCreateExample{}
	e.Init()

	requestJsonDoc := strings.Replace(`
		{
			"advertiser_id":{advertiser_id},
			"operation":"ENABLE",
			"delivery_mode":"MANUAL",
			"landing_type":"LINK",
			"marketing_goal":"VIDEO_AND_IMAGE",
			"ad_type":"ALL",
			"name":"sdk测试4",
			"search_bid_ratio":1.05,
			"audience_extend":"ON",
			"related_product":{
				"product_setting":"SINGLE",
				"product_platform_id":1805155490137893,
				"product_id":"1688630940926924743"
			},
			"promotion_type":"LANDING_PAGE_LINK",
			"asset_type":"ORANGE",
			"dpa_adtype":"DPA_LINK",
			"optimize_goal":{
				"external_action": "AD_CONVERT_TYPE_PAY"
			},
			"value_optimized_type":"OFF",
			"landing_page_stay_time":5000,
			"delivery_range":{
				"inventory_catalog":"UNIVERSAL_SMART"
			},
			"audience":{
				"district": "REGION",
                "age": [
                    "AGE_BETWEEN_24_30",
                    "AGE_BETWEEN_31_40",
                    "AGE_BETWEEN_41_49",
                    "AGE_ABOVE_50"
                ],
				"region_version": "1.0.0",
				"city": [34,50,35,44,62,45,52,13,42,23,46,41,43,22,32,36,21,15,64,63,51,37,31,61,14,12,65,54,53,33],
				"location_type": "CURRENT",
				"interest_action_mode": "UNLIMITED",
				"ac": ["WIFI","4G","5G"],
				"hide_if_exists": "UNLIMITED",
				"auto_extend_targets": [
                    "AGE"
                ]
			},
			"delivery_setting": {
                "schedule_type": "SCHEDULE_FROM_NOW",
                "schedule_time": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
                "bid_type": "CUSTOM",
                "bid_speed": "FAST",
                "budget_mode": "BUDGET_MODE_INFINITE"
            },
			"track_url_setting": {
                "send_type": "SERVER_SEND"
            }
		}
	`, "{advertiser_id}", AdvertiserId, -1)

	json.Unmarshal([]byte(requestJsonDoc), &e.ProjectCreateRequest)

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
