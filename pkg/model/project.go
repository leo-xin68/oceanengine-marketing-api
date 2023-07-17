package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
)

// Keywords 搜索关键词列表
type Keywords struct {
	Word      *string   `json:"word,omitempty"`     //关键词
	BidType   *string   `json:"bid_type,omitempty"` //出价类型。 允许值:FEED_TO_SEARCH 搜索快投
	MatchType MatchType `json:"match_type,omitempty"`
}

// RelatedProduct 关联产品投放相关
type RelatedProduct struct {
	ProductSetting    *string     `json:"product_setting,omitempty"`     //商品库设置，枚举值：SINGLE 启用SDPA、NO_MAP不启用
	ProductPlatformId *uint64     `json:"product_platform_id,omitempty"` //商品库ID
	ProductId         *string     `json:"product_id,omitempty"`          //产品ID
	Products          *[]Products `json:"products,omitempty"`            //产品ID列表
}

// Products 产品ID列表
type Products struct {
	ProductId         *string `json:"product_id,omitempty"`          //产品ID
	ProductPlatformId *uint64 `json:"product_platform_id,omitempty"` //产品库ID
	AssetId           *string `json:"asset_id,omitempty"`            //物件ID
}

// CreateOptimizeGoal 创建优化目标
type CreateOptimizeGoal struct {
	AssetIds           *[]uint64 `json:"asset_ids,omitempty"`            //事件管理资产id
	ConvertId          *uint64   `json:"convert_id,omitempty"`           //转化跟踪id
	ExternalAction     *string   `json:"external_action,omitempty"`      //优化目标
	DeepExternalAction *string   `json:"deep_external_action,omitempty"` //深度转化目标
	PaidSwitch         *uint8    `json:"paid_switch,omitempty"`          //字节提供的归因方式，返回值： 1启用；2 不启用
}

// OptimizeGoal 优化目标
type OptimizeGoal struct {
	CreateOptimizeGoal
	ValueOptimizedType *OptimizedType `json:"value_optimized_type,omitempty"` //目标优化类型
}

// DeliveryRange 广告版位
type DeliveryRange struct {
	InventoryCatalog InventoryCatalog `json:"inventory_catalog,omitempty"`
	InventoryType    []InventoryType  `json:"inventory_type,omitempty"`
	UnionVideoType   UnionVideoType   `json:"union_video_type,omitempty"`
}

// Geolocation 地图位置
type Geolocation struct {
	Radius *uint64  `json:"radius,omitempty"` //半径，单位为m
	Name   *string  `json:"name,omitempty"`   //地点名称
	Long   *float64 `json:"long,omitempty"`   //经度
	Lat    *float64 `json:"lat,omitempty"`    //纬度
}

// DeliverySetting 投放设置
type DeliverySetting struct {
	ScheduleType         *ScheduleType `json:"schedule_type,omitempty"`
	StartTime            *string       `json:"start_time,omitempty"`          //投放起始时间，如：2017-01-01 精确到天
	EndTime              *string       `json:"end_time,omitempty"`            //投放结束时间，如：2017-01-01 精确到天
	ScheduleTime         *string       `json:"schedule_time,omitempty"`       //投放时段
	FilterNightSwitch    *Switch       `json:"filter_night_switch,omitempty"` //过滤夜间投放
	DeepBidType          *DeepBidType  `json:"deep_bid_type,omitempty"`
	BidType              *string       `json:"bid_type,omitempty"`       //竞价策略，枚举值：CUSTOM 稳定成本、NO_BID 最大转化投放、UPPER_CONTROL最优成本
	ProjectCustom        *Switch       `json:"project_custom,omitempty"` //项目成本稳投
	Bid                  *float32      `json:"bid,omitempty"`            //点击出价/展示出价
	BidSpeed             *BidSpeed     `json:"bid_speed,omitempty"`
	BudgetMode           *string       `json:"budget_mode,omitempty"`            //项目预算类型， 枚举值：BUDGET_MODE_INFINITE 不限、BUDGET_MODE_DAY 日预算
	Budget               *float32      `json:"budget,omitempty"`                 //项目预算
	CpaBid               *float32      `json:"cpa_bid,omitempty"`                //目标转化出价/预期成本(注意：nobid不返回该字段)
	DeepCpabid           *float32      `json:"deep_cpabid,omitempty"`            //深度优化出价
	RoiGoal              *float32      `json:"roi_goal,omitempty"`               //深度转化ROI系数(注意：nobid不返回该字段)
	BudgetOptimizeSwitch *Switch       `json:"budget_optimize_switch,omitempty"` //支持预算择优分配，枚举值： ON 开启，OFF 不开启
}

// CreateDeliverySetting 创建投放设置 (排期、预算与出价)
type CreateDeliverySetting struct {
	Pricing *PricingType `json:"pricing,omitempty"`
	DeliverySetting
}

// TrackUrlSetting 监测链接设置
type TrackUrlSetting struct {
	TrackUrlType               *TrackUrlType `json:"track_url_type,omitempty"`                 //监测链接类型，区分使用监测链接组或者自定义链接 枚举值：CUSTOM 自定义链接、GROUP_ID 监测链接组
	TrackUrlGroupId            *uint64       `json:"track_url_group_id,omitempty"`             //监测链接组id
	TrackUrl                   *[]string     `json:"track_url,omitempty"`                      //展示（监测链接）
	ActionTrackUrl             *[]string     `json:"action_track_url,omitempty"`               //点击（监测链接）
	VideoPlayFirstTrackUrl     *[]string     `json:"video_play_first_track_url,omitempty"`     //视频开始播放（监测链接）
	VideoPlayDoneTrackUrl      *[]string     `json:"video_play_done_track_url,omitempty"`      //视频播完（监测链接）
	VideoPlayEffectiveTrackUrl *[]string     `json:"video_play_effective_track_url,omitempty"` //视频有效播放（监测链接）
	SendType                   *SendType     `json:"send_type,omitempty"`
}

// DpaProductTarget 自定义筛选条件（商品投放条件）
type DpaProductTarget struct {
	Title *[]string `json:"title,omitempty"` //筛选字段
	Rule  *[]string `json:"rule,omitempty"`  //定向规则
	Type  *[]string `json:"type,omitempty"`  //字段类型
	Value *[]string `json:"value,omitempty"` //规则值
}

// Audience 定向 (商品定向跟人群定向合一起)
type Audience struct {
	//人群定向（线索智投场景）
	AudiencePackageId      *uint64        `json:"audience_package_id,omitempty"`      //定向包ID 如果同时传定向包ID和自定义用户定向参数时，仅定向包中的定向生效
	District               *string        `json:"district,omitempty"`                 //地域类型，枚举值: CITY 省市、 COUNTY 区县、BUSINESS_DISTRICT 商圈、REGION 行政区域、OVERSEA 海外区域、NONE 不限
	Geolocation            *[]Geolocation `json:"geolocation,omitempty"`              //从地图添加(地图位置)
	RegionRecommend        *string        `json:"region_recommend,omitempty"`         //地域智能放量定向，ON为开启、OFF为关闭
	Age                    *[]AgeInterval `json:"age,omitempty"`                      //年龄
	RetargetingTagsInclude *[]uint64      `json:"retargeting_tags_include,omitempty"` //定向人群包列表（自定义人群）
	RetargetingTagsExclude *[]uint64      `json:"retargeting_tags_exclude,omitempty"` //排除人群包列表（自定义人群）
	HideIfConverted        *string        `json:"hide_if_converted,omitempty"`        //过滤已转化用户 枚举值：NO_EXCLUDE 不限制、PROMOTION 广告、PROJECT 推广项目、ADVERTISER 广告账户、APP 应用、CUSTOMER 客户、ORGANIZATION 组织
	//人群定向
	RegionVersion             *string                `json:"region_version,omitempty"` //行政区域版本号
	City                      *[]uint64              `json:"city,omitempty"`           //地域定向省市或者区县列表，
	LocationType              *LocationType          `json:"location_type,omitempty"`
	Gender                    Gender                 `json:"gender,omitempty"`                   //性别
	InterestActionMode        *string                `json:"interest_action_mode,omitempty"`     //行为兴趣，枚举值：UNLIMITED 不限、CUSTOM 自定义、 RECOMMEND系统推荐
	ActionDays                *uint64                `json:"action_days,omitempty"`              //用户发生行为天数，枚举值：7、 15、 30、 60、 90、 180、 365用户发生行为天数
	ActionCategories          *[]uint64              `json:"action_categories,omitempty"`        //行为类目词
	ActionWords               *[]uint64              `json:"action_words,omitempty"`             //行为关键词
	InterestCategories        *[]uint64              `json:"interest_categories,omitempty"`      //兴趣类目词
	InterestWords             *[]uint64              `json:"interest_words,omitempty"`           //兴趣关键词
	AwemeFanBehaviors         *[]string              `json:"aweme_fan_behaviors,omitempty"`      //抖音达人互动用户行为类型
	AwemeFanTimeScope         *string                `json:"aweme_fan_time_scope,omitempty"`     //抖音达人互动行为时间范围，枚举值：FIFTEEN_DAYS 15天、THIRTY_DAYS 30天、SIXTY_DAYS 60天
	AwemeFanCategories        *[]uint64              `json:"aweme_fan_categories,omitempty"`     //抖音达人分类ID列表
	AwemeFanAccounts          *[]uint64              `json:"aweme_fan_accounts,omitempty"`       //抖音达人ID列表
	SuperiorPopularityType    *string                `json:"superior_popularity_type,omitempty"` //媒体定向
	ExcludeFlowPackage        *[]uint64              `json:"exclude_flow_package,omitempty"`     //排除定向逻辑
	Platform                  *[]string              `json:"platform,omitempty"`                 //投放平台列表,枚举值：ANDROID、IOS
	AndroidOsv                *string                `json:"android_osv,omitempty"`              //最低安卓版本
	IosOsv                    *string                `json:"ios_osv,omitempty"`                  //最低IOS版本
	DeviceType                *[]string              `json:"device_type,omitempty"`              //设备类型，枚举值：MOBILE、PAD
	Ac                        *[]AcType              `json:"ac,omitempty"`
	Carrier                   *[]CarrierType         `json:"carrier,omitempty"`
	HideIfExists              *string                `json:"hide_if_exists,omitempty"`               //过滤已安装，枚举值：UNLIMITED不限、FILTER 过滤、TARGETING 定向
	ConvertedTimeDuration     *ConvertedTimeDuration `json:"converted_time_duration,omitempty"`      //过滤时间范围
	FilterAwemeAbnormalActive *string                `json:"filter_aweme_abnormal_active,omitempty"` //过滤高活跃用户，即过滤关注、点赞、评论行为高活跃的用户允许值： ON 过滤 OFF不过滤（默认值）
	FilterAwemeFansCount      *uint64                `json:"filter_aweme_fans_count,omitempty"`      //过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户 允许值：1000、500、200
	FilterOwnAwemeFans        *string                `json:"filter_own_aweme_fans,omitempty"`        //过滤自己的粉丝，允许值：ON 过滤 OFF不过滤（默认值）
	DeviceBrand               *[]DeviceBrand         `json:"device_brand,omitempty"`                 //手机品牌
	LaunchPrice               *[]uint64              `json:"launch_price,omitempty"`                 //手机价格
	AutoExtendTargets         *[]string              `json:"auto_extend_targets,omitempty"`          //可放开定向，枚举值：AGE 年龄、REGION 地域、GENDER 性别、CUSTOM_AUDIENCE 自定人群-定向
	// 商品定向
	DpaCity             *Switch   `json:"dpa_city,omitempty"`               //地域匹配-商品所在城市（OFF表示不启用，ON表示启用）
	DpaLbs              *Switch   `json:"dpa_lbs,omitempty"`                //地域匹配-适地性服务开启时，根据用户的地理位置信息，给用户投放位于其附近的产品
	DpaRtaSwitch        *Switch   `json:"dpa_rta_switch,omitempty"`         //RTA重定向开关（OFF表示不启用，ON表示启用）
	FlowPackage         *[]uint64 `json:"flow_package,omitempty"`           //定向逻辑
	RtaId               *[]uint64 `json:"rta_id,omitempty"`                 //RTA策略ID
	DpaRtaRecommendType *string   `json:"dpa_rta_recommend_type,omitempty"` //RTA推荐逻辑

	ActionScene *[]string `json:"action_scene,omitempty"` //行为场景
}
