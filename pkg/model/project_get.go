package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type ProjectGetRequest struct {
	AdvertiserId *uint64                     `json:"advertiser_id,omitempty"`
	Fields       []*string                   `json:"fields,omitempty"`
	Filtering    *ProjectGetRequestFiltering `json:"filtering,omitempty"`
	Page         *uint64                     `json:"page,omitempty"`
	PageSize     *uint64                     `json:"page_size,omitempty"`
}

type ProjectGetRequestFiltering struct {
	IDs               *[]uint64             `json:"ids,omitempty"`
	DeliveryMode      DeliveryMode          `json:"delivery_mode,omitempty"`
	LandingType       LandingType           `json:"landing_type,omitempty"`
	AppPromotionType  AppPromotionType      `json:"app_promotion_type,omitempty"`
	MarketingGoal     MarketingGoal         `json:"marketing_goal,omitempty"`
	AdType            AdType                `json:"ad_type,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Status            ProjectStatus         `json:"status,omitempty"`
	StatusFirst       ProjectFirstStatus    `json:"status_first,omitempty"`
	StatusSecond      []ProjectSecondStatus `json:"status_second,omitempty"`
	ProjectCreateTime *string               `json:"project_create_time,omitempty"`
	ProjectModifyTime *string               `json:"project_modify_time,omitempty"`
	Pricing           *PricingType          `json:"pricing,omitempty"`
	InventoryType     *InventoryType        `json:"inventory_type,omitempty"`
	Platform          *string               `json:"platform,omitempty"`
}

type ProjectGetResponse struct {
	common.ResponseBaseInfo
	Data *ProjectGetResponseData `json:"data,omitempty"`
}

type ProjectGetResponseData struct {
	List     []ProjectGetListStruct   `json:"list,omitempty"`
	PageInfo *common.ResponsePageInfo `json:"page_info,omitempty"`
}

type ProjectGetListStruct struct {
	ProjectId           *uint64               `json:"project_id,omitempty"`    //项目ID
	AdvertiserId        *uint64               `json:"advertiser_id,omitempty"` //广告账户id
	DeliveryMode        DeliveryMode          `json:"delivery_mode,omitempty"`
	LandingType         LandingType           `json:"landing_type,omitempty"`
	AppPromotionType    AppPromotionType      `json:"app_promotion_type,omitempty"`
	MarketingGoal       MarketingGoal         `json:"marketing_goal,omitempty"`
	AdType              AdType                `json:"ad_type,omitempty"`
	OptStatus           *string               `json:"opt_status,omitempty"`          //目标操作，枚举值：ENABLE 启用项目、DISABLE暂停项目。
	Name                *string               `json:"name,omitempty"`                //项目名称
	ProjectCreateTime   *string               `json:"project_create_time,omitempty"` //项目创建时间，格式yyyy-MM-dd HH:mm:ss
	ProjectModifyTime   *string               `json:"project_modify_time,omitempty"` //项目更新时间，格式yyyy-MM-dd HH:mm:ss
	Status              ProjectStatus         `json:"status,omitempty"`
	StatusFirst         ProjectFirstStatus    `json:"status_first,omitempty"`
	StatusSecond        []ProjectSecondStatus `json:"status_second,omitempty"`
	Pricing             PricingType           `json:"pricing,omitempty"`
	PackageName         *string               `json:"package_name,omitempty"`         //应用包名
	AppName             *string               `json:"app_name,omitempty"`             //应用名
	FeedDeliverySearch  *string               `json:"feed_delivery_search,omitempty"` // 搜索快投关键词功能，HAS_OPEN:启用，DISABLED:未启用
	SearchBidRatio      *float32              `json:"search_bid_ratio,omitempty"`     //出价系数
	AudienceExtend      Switch                `json:"audience_extend,omitempty"`      //定向拓展
	Keywords            *[]Keywords           `json:"keywords,omitempty"`
	RelatedProduct      *RelatedProduct       `json:"related_product,omitempty"`
	Title               *string               `json:"title,omitempty"` //筛选字段
	Rule                *string               `json:"rule,omitempty"`  //定向规则
	Type                *string               `json:"type,omitempty"`  //字段类型
	Value               *string               `json:"value,omitempty"` //规则值
	DpaAdType           DpaAdType             `json:"dpa_adtype,omitempty"`
	DownloadUrl         *string               `json:"download_url,omitempty"`  //下载链接
	DownloadType        *string               `json:"download_type,omitempty"` //下载方式，枚举值：DOWNLOAD_URL 直接下载、EXTERNAL_URL 落地页下载
	DownloadMode        DownloadMode          `json:"download_mode,omitempty"`
	LaunchType          LaunchType            `json:"launch_type,omitempty"`
	PromotionType       PromotionType         `json:"promotion_type,omitempty"`
	OpenUrl             *string               `json:"open_url,omitempty"`      //Deeplink直达链接
	UlinkUrl            *string               `json:"ulink_url,omitempty"`     //ulink直达链接
	SubscribeUrl        *string               `json:"subscribe_url,omitempty"` //预约下载链接
	AssetType           AssetType             `json:"asset_type,omitempty"`
	OptimizeGoal        *OptimizeGoal         `json:"optimize_goal,omitempty"`          //优化目标
	LandingPageStayTime *uint64               `json:"landing_page_stay_time,omitempty"` //店铺停留时长，单位为毫秒
	DeliveryRange       *DeliveryRange        `json:"delivery_range,omitempty"`         //广告版位
	Audience            *Audience             `json:"audience,omitempty"`               //定向设置
	DeliverySetting     *DeliverySetting      `json:"delivery_setting,omitempty"`       //投放设置
	TrackUrlSetting     *TrackUrlSetting      `json:"track_url_setting,omitempty"`      //监测链接设置
	OpenUrlType         *string               `json:"open_url_type,omitempty"`          //直达链接类型
	OpenUrlField        *string               `json:"open_url_field,omitempty"`         //直达链接字段选择
	OpenUrlParams       *string               `json:"open_url_params,omitempty"`        //直达链接检测参数
}
