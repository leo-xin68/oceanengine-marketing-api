package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type PromotionGetRequest struct {
	AdvertiserId *uint64                       `json:"advertiser_id,omitempty"`
	Fields       []*string                     `json:"fields,omitempty"`
	Filtering    *PromotionGetRequestFiltering `json:"filtering,omitempty"`
	Page         *uint64                       `json:"page,omitempty"`
	PageSize     *uint64                       `json:"page_size,omitempty"`
}

type PromotionGetRequestFiltering struct {
	IDs                 *[]uint64             `json:"ids,omitempty"`
	Name                *string               `json:"name,omitempty"`
	ProjectId           *uint64               `json:"project_id,omitempty"`
	DeliveryMode        DeliveryMode          `json:"delivery_mode,omitempty"`
	Status              PromotionStatus       `json:"status,omitempty"`
	StatusFirst         PromotionFirstStatus  `json:"status_first,omitempty"`
	StatusSecond        PromotionSecondStatus `json:"status_second,omitempty"`
	PromotionCreateTime *string               `json:"promotion_create_time,omitempty"`
	PromotionModifyTime *string               `json:"promotion_modify_time,omitempty"`
	RejectReasonType    RejectReasonType      `json:"reject_reason_type,omitempty"`
	LearningPhase       *[]string             `json:"learning_phase,omitempty"`
}

type PromotionGetResponse struct {
	common.ResponseBaseInfo
	Data *PromotionGetResponseData `json:"data,omitempty"`
}

type PromotionGetResponseData struct {
	List     []PromotionGetListStruct `json:"list,omitempty"`
	PageInfo *common.ResponsePageInfo `json:"page_info,omitempty"`
}

type PromotionGetListStruct struct {
	PromotionId         *uint64                 `json:"promotion_id,omitempty"`          //广告ID
	PromotionName       *string                 `json:"promotion_name,omitempty"`        //广告名称
	ProjectId           *uint64                 `json:"project_id,omitempty"`            //项目ID
	AdvertiserId        *uint64                 `json:"advertiser_id,omitempty"`         //广告账户id
	PromotionCreateTime *string                 `json:"promotion_create_time,omitempty"` //项目创建时间，格式yyyy-MM-dd HH:mm:ss
	PromotionModifyTime *string                 `json:"promotion_modify_time,omitempty"` //项目更新时间，格式yyyy-MM-dd HH:mm:ss
	LearningPhase       LearningPhase           `json:"learning_phase,omitempty"`
	Status              PromotionStatus         `json:"status,omitempty"`
	StatusFirst         PromotionFirstStatus    `json:"status_first,omitempty"`
	StatusSecond        []PromotionSecondStatus `json:"status_second,omitempty"`
	OptStatus           *string                 `json:"opt_status,omitempty"` // 操作状态
	NativeSetting       *NativeSetting          `json:"native_setting,omitempty"`
	PromotionMaterials  *PromotionMaterials     `json:"promotion_materials,omitempty"`
	Source              *string                 `json:"source,omitempty"`              //广告来源
	IsCommentDisable    *Switch                 `json:"is_comment_disable,omitempty"`  //广告评论
	AdDownloadStatus    *Switch                 `json:"ad_download_status,omitempty"`  //客户端下载视频功能
	MaterialsType       *string                 `json:"materials_type,omitempty"`      //素材类型
	Budget              *float64                `json:"budget,omitempty"`              //预算
	BudgetMode          *string                 `json:"budget_mode,omitempty"`         //预算类型
	Bid                 *uint64                 `json:"bid,omitempty"`                 //点击出价/展示出价
	CpaBid              *float64                `json:"cpa_bid,omitempty"`             //目标转化出价/预期成本
	DeepCpabid          *uint64                 `json:"deep_cpabid,omitempty"`         //深度优化出价
	RoiGoal             *uint64                 `json:"roi_goal,omitempty"`            //深度转化ROI系数
	ScheduleTime        *string                 `json:"schedule_time,omitempty"`       //广告的投放时段
	MaterialScoreInfo   *MaterialScoreInfo      `json:"material_score_info,omitempty"` //素材评级信息
}
