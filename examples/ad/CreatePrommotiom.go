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

type PromotionCreateExample struct {
	OceanAds               *ads.App
	PromotionCreateRequest model.PromotionCreateRequest
}

func (e *PromotionCreateExample) Init() {
	cfg := config.NewSdkConfig(AppId, AccessToken, true)
	e.OceanAds = ads.Init(cfg)
}

func (e *PromotionCreateExample) RunExample() (model.PromotionCreateResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Promotion.Create(ctx, e.PromotionCreateRequest)
}

func main() {
	e := &PromotionCreateExample{}
	e.Init()

	requestJsonDoc := strings.Replace(`
		{
			"advertiser_id": {advertiser_id},
			"project_id": 7254855571727089683,
			"name": "sdk测试",
			"native_setting": {
				"aweme_id": "98208682798",
				"is_feed_and_fav_see": "NO",
				"anchor_related_type": "SELECT"
			},
			"promotion_materials": {
				"video_material_list": [
					{
						"image_mode": "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
						"video_id": "v03033g10000cinujhjc77uckg24nl6g",
						"video_cover_id": "tos-cn-p-0051/ogfBFLjNOvAkfAAMETgNN61zghrDDkBApRrjIZ"
					}
				],
				"title_material_list": [
					{
						"title": "替姐姐嫁给残疾少爷，以为是跳入深渊，没想到婚后被宠上天！"
					}
				],
				"external_url_material_list": [
					"https://www.chengzijianzhan.cc/tetris/page/7255257569093746749?projectid=__PROJECT_ID__&promotionid=__PROMOTION_ID__&creativetype=__CTYPE__&clickid=__CLICKID__&mid1=__MID1__&mid2=__MID2__&mid3=__MID3__&mid4=__MID4__&mid5=__MID5__"
				],
				"product_info": {
					"titles": [
						"苏小姐，你嫁错人了"
					],
					"image_ids": [
						"web.business.image/202307135d0d10fbee6891e84b84864c"
					],
					"selling_points": [
						"年度热门短剧"
					]
				},
				"call_to_action_buttons": [
					"点击观看"
				],
				"anchor_material_list": [
					{
						"anchor_type": "APP_INTERNET_SERVICE",
						"anchor_id": "7255238507369696041"
					}
				],
				"dynamic_creative_switch": "OFF",
				"intelligent_generation": "ON"
			},
			"source": "点锐剧场",
			"is_comment_disable": "OFF",
			"budget": 9999999.99,
			"budget_mode": "BUDGET_MODE_DAY",
			"cpa_bid": 1,
			"operation": "DISABLE"
		}
	`, "{advertiser_id}", AdvertiserId, -1)

	json.Unmarshal([]byte(requestJsonDoc), &e.PromotionCreateRequest)

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
