package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type ProjectCreateRequest struct {
	AdvertiserId        *uint64                `json:"advertiser_id,omitempty"`
	Name                *string                `json:"name,omitempty"`
	LandingType         LandingType            `json:"landing_type,omitempty"`
	AppPromotionType    AppPromotionType       `json:"app_promotion_type,omitempty"`
	MarketingGoal       MarketingGoal          `json:"marketing_goal,omitempty"`
	AdType              AdType                 `json:"ad_type,omitempty"`
	LaunchType          LaunchType             `json:"launch_type,omitempty"`
	OpenUrl             *string                `json:"open_url,omitempty"`      //Deeplink直达链接
	UlinkUrl            *string                `json:"ulink_url,omitempty"`     //ulink直达链接
	SubscribeUrl        *string                `json:"subscribe_url,omitempty"` //预约下载链接
	DownloadType        DownloadType           `json:"download_type,omitempty"`
	DownloadUrl         *string                `json:"download_url,omitempty"` //下载链接
	DownloadMode        DownloadMode           `json:"download_mode,omitempty"`
	AssetType           AssetType              `json:"asset_type,omitempty"`
	MicroPromotionType  MicroPromotionType     `json:"micro_promotion_type,omitempty"`
	LandingPageStayTime LandingPageStayTime    `json:"landing_page_stay_time,omitempty"`
	QuickAppId          *uint64                `json:"quick_app_id,omitempty"` //快应用资产id
	RelatedProduct      *RelatedProduct        `json:"related_product,omitempty"`
	PromotionType       PromotionType          `json:"promotion_type,omitempty"`
	DpaCategories       *[]uint64              `json:"dpa_categories,omitempty"` //商品投放范围，分类列表
	DpaProductTarget    *DpaProductTarget      `json:"dpa_product_target,omitempty"`
	OptimizeGoal        *CreateOptimizeGoal    `json:"optimize_goal,omitempty"`  //优化目标
	DeliveryRange       *DeliveryRange         `json:"delivery_range,omitempty"` //广告版位
	DeliverySetting     *CreateDeliverySetting `json:"delivery_setting,omitempty"`
	TrackUrlSetting     *TrackUrlSetting       `json:"track_url_setting,omitempty"`
	Audience            *Audience              `json:"audience,omitempty"` //定向设置
	Operation           *string                `json:"operation,omitempty"`
	DeliveryMode        DeliveryMode           `json:"delivery_mode,omitempty"`
	SearchBidRatio      *float32               `json:"search_bid_ratio,omitempty"`
	AudienceExtend      Switch                 `json:"audience_extend,omitempty"`
	Keywords            *[]Keywords            `json:"keywords,omitempty"`
	DpaAdType           DpaAdType              `json:"dpa_adtype,omitempty"`
	ValueOptimizedType  OptimizedType          `json:"value_optimized_type,omitempty"` //目标优化类型
	OpenUrlType         *string                `json:"open_url_type,omitempty"`        //直达链接类型
	OpenUrlField        *string                `json:"open_url_field,omitempty"`       //直达链接字段选择
	OpenUrlParams       *string                `json:"open_url_params,omitempty"`      //直达链接检测参数
}

type ProjectCreateResponse struct {
	common.ResponseBaseInfo
	Data *ProjectCreateResponseData `json:"data,omitempty"`
}

type ProjectCreateResponseData struct {
	ProjectId *uint64 `json:"project_id,omitempty"` //项目ID
}
