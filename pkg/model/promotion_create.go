package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type PromotionCreateRequest struct {
	AdvertiserId               *uint64             `json:"advertiser_id,omitempty"`
	ProjectId                  *uint64             `json:"project_id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	PromotionMaterials         *PromotionMaterials `json:"promotion_materials,omitempty"`
	Budget                     *float64            `json:"budget,omitempty"`      //预算
	BudgetMode                 BudgetMode          `json:"budget_mode,omitempty"` //预算类型
	CpaBid                     *float64            `json:"cpa_bid,omitempty"`     //目标转化出价/预期成本
	Bid                        *uint64             `json:"bid,omitempty"`         //点击出价/展示出价
	DeepCpabid                 *uint64             `json:"deep_cpabid,omitempty"` //深度优化出价
	RoiGoal                    *uint64             `json:"roi_goal,omitempty"`    //深度转化ROI系数
	Source                     *string             `json:"source,omitempty"`      // 广告来源
	Operation                  *string             `json:"operation,omitempty"`
	NativeSetting              *NativeSetting      `json:"native_setting,omitempty"`
	IsCommentDisable           Switch              `json:"is_comment_disable,omitempty"` //广告评论
	AdDownloadStatus           Switch              `json:"ad_download_status,omitempty"` //客户端下载视频功能
	MaterialsType              MaterialsType       `json:"materials_type,omitempty"`
	CreativeAutoGenerateSwitch Switch              `json:"creative_auto_generate_switch,omitempty"` //是否开启自动生成素材
	ConfigId                   *uint64             `json:"config_id,omitempty"`                     //配置ID
	AutoExtendTraffic          Switch              `json:"auto_extend_traffic,omitempty"`           //智能拓流
	Keywords                   *PromotionKeywords  `json:"keywords,omitempty"`
	BrandInfo                  *BrandInfo          `json:"brand_info,omitempty"` //品牌信息
}

type PromotionCreateResponse struct {
	common.ResponseBaseInfo
	Data *PromotionCreateResponseData `json:"data,omitempty"`
}

type PromotionCreateResponseData struct {
	PromotionId       *uint64        `json:"promotion_id,omitempty"`        //广告ID
	ErrorKeywordsList *ErrorKeywords `json:"error_keywords_list,omitempty"` //广告ID
}

// ErrorKeywords 设置失败的搜索广告关键词
type ErrorKeywords struct {
	ErrorKeyword *string `json:"error_keyword,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
}
