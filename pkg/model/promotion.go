package model

import (
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/enums"
)

// PromotionMaterials 广告素材组合
type PromotionMaterials struct {
	VideoMaterialList       *[]VideoMaterial      `json:"video_material_list,omitempty"`
	ImageMaterialList       *[]ImageMaterial      `json:"image_material_list,omitempty"`
	TitleMaterialList       *[]TitleMaterial      `json:"title_material_list,omitempty"`
	TextAbstractList        *TextAbstract         `json:"text_abstract_list,omitempty"`
	ParamsType              *string               `json:"params_type,omitempty"`                //链接类型(落地页)
	ExternalUrlField        *string               `json:"external_url_field,omitempty"`         //落地页链接字段选择
	ExternalUrlMaterialList *[]string             `json:"external_url_material_list,omitempty"` // 普通落地页链接素材
	ExternalUrlParams       *string               `json:"external_url_params,omitempty"`        //落地页检测参数
	WebUrlMaterialList      *[]string             `json:"web_url_material_list,omitempty"`      // Android应用下载详情页
	PlayableUrlMaterialList *[]string             `json:"playable_url_material_list,omitempty"` // 试玩落地页素材
	ProductInfo             *ProductInfo          `json:"product_info,omitempty"`
	CallToActionButtons     *[]string             `json:"call_to_action_buttons,omitempty"` // 行动号召文案
	ComponentMaterialList   *[]ComponentMaterial  `json:"component_material_list,omitempty"`
	AnchorMaterialList      *[]AnchorMaterial     `json:"anchor_material_list,omitempty"`
	MiniProgramInfo         *MiniProgramInfo      `json:"mini_program_info,omitempty"`
	OpenUrl                 *string               `json:"open_url,omitempty"`                //直达链接，用于打开电商app，调起店铺
	OpenUrlType             *string               `json:"open_url_type,omitempty"`           //直达链接类型
	OpenUrlField            *string               `json:"open_url_field,omitempty"`          //直达链接字段选择
	OpenUrlParams           *string               `json:"open_url_params,omitempty"`         //直达链接检测参数
	Ulink                   *string               `json:"ulink,omitempty"`                   //直达备用链接
	DynamicCreativeSwitch   Switch                `json:"dynamic_creative_switch,omitempty"` // 动态创意开关
	IntelligentGeneration   *string               `json:"intelligent_generation,omitempty"`  // 智能生成行动号召按钮
	DecorationMaterial      *[]DecorationMaterial `json:"decoration_material,omitempty"`
}

// VideoMaterial 视频素材信息
type VideoMaterial struct {
	ImageMode         *string   `json:"image_mode,omitempty"`          //素材类型
	VideoId           *string   `json:"video_id,omitempty"`            //视频ID
	VideoCoverId      *string   `json:"video_cover_id,omitempty"`      //视频封面图片ID
	ItemId            *uint64   `json:"item_id,omitempty"`             //抖音短视频ID
	VideoTemplateType *string   `json:"video_template_type,omitempty"` //商品库视频生成类型
	VideoTaskIds      *[]string `json:"video_task_ids,omitempty"`      //自定义商品库视频模板ID
	Images            *[]Image  `json:"images,omitempty"`

	MaterialId     *uint64        `json:"material_id,omitempty"` //素材ID
	MaterialStatus MaterialStatus `json:"material_status,omitempty"`
}

// ImageMaterial 创意图片素材
type ImageMaterial struct {
	ImageMode *string  `json:"image_mode,omitempty"` //素材类型
	Images    *[]Image `json:"images,omitempty"`     //图片ID数组
}

// Image 图片
type Image struct {
	ImageId          *uint64         `json:"image_id,omitempty"`    //图片ID
	TemplateId       *uint64         `json:"template_id,omitempty"` //图片素材类型-DPA模板ID
	TemplateDataList *[]TemplateData `json:"template_data_list,omitempty"`

	MaterialId     *uint64        `json:"material_id,omitempty"` //素材ID
	MaterialStatus MaterialStatus `json:"material_status,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	Title       *string    `json:"title,omitempty"`
	WordList    *[]uint64  `json:" word_list,omitempty"`
	BidWordList *[]BidWord `json:"bidword_list,omitempty"`  //搜索关键词列表
	DpaWordList *[]uint64  `json:"dpa_word_list,omitempty"` //DPA词包ID列表

	MaterialId *uint64 `json:"material_id,omitempty"`
}

// BidWord 搜索关键词
type BidWord struct {
	DefaultWord *string `json:"default_word,omitempty"` //关键词
}

// TextAbstract 文本摘要信息
type TextAbstract struct {
	AbstractText *string    `json:"abstract_text,omitempty"` //文本摘要内容
	BidwordList  *[]BidWord `json:"bidword_list,omitempty"`
	WordList     *[]uint64  `json:"word_list,omitempty"`
}

// DecorationMaterial 家装卡券素材
type DecorationMaterial struct {
	ImageMode  *string `json:"image_mode,omitempty"`  //素材类型
	ActivityId *string `json:"activity_id,omitempty"` //活动ID

	MaterialId *uint64 `json:"material_id,omitempty"` //素材ID
}

// WordId 动态词包ID
type WordId struct {
	WordId *uint64 `json:"word_id,omitempty"` //关键词
}

// ComponentMaterial 创意组件
type ComponentMaterial struct {
	ComponentId *uint64 `json:"component_id,omitempty"` //组件id
}

// AnchorMaterial 锚点信息
type AnchorMaterial struct {
	AnchorId   *string    `json:"anchor_id,omitempty"`   //锚点id
	AnchorType AnchorType `json:"anchor_type,omitempty"` //锚点类型
}

// MiniProgramInfo 字节小程序信息
type MiniProgramInfo struct {
	AppId     *string `json:"app_id,omitempty"`     //小程序/小游戏id
	StartPath *string `json:"start_path,omitempty"` //启动路径
	Params    *string `json:"params,omitempty"`     //页面监测参数
	Url       *string `json:"url,omitempty"`        //字节小程序调起链接
}

// ProductInfo 产品信息
type ProductInfo struct {
	Titles                    *[]string `json:"titles,omitempty"`                       //产品名称
	ImageIds                  *[]string `json:"image_ids,omitempty"`                    //产品主图
	SellingPoints             *[]string `json:"selling_points,omitempty"`               //产品卖点
	ProductNameType           *string   `json:"product_name_type,omitempty"`            //产品名称类型
	ProductImageType          *string   `json:"product_image_type,omitempty"`           //产品图片类型
	ProductSellingPointType   *string   `json:"product_selling_point_type,omitempty"`   //产品卖点类型
	ProductNameFields         *[]string `json:"product_name_fields,omitempty"`          //DPA产品库名称字段
	ProductImageFields        *[]string `json:"product_image_fields,omitempty"`         //DPA产品库图片字段
	ProductSellingPointFields *[]string `json:"product_selling_point_fields,omitempty"` //DPA产品库卖点类型字段
}

// MaterialScoreInfo 素材评级信息
type MaterialScoreInfo struct {
	ScoreNumOfMaterial     *string               `json:"score_num_of_material,omitempty"`     //素材数量评级分数
	ScoreTypeOfMaterial    *string               `json:"score_type_of_material,omitempty"`    //素材类型评级分数
	ScoreValueOfMaterial   *string               `json:"score_value_of_material,omitempty"`   //素材品质评级分数
	MaterialAdvice         *string               `json:"material_advice,omitempty"`           //素材评级建议
	LowQualityMaterialList *[]LowQualityMaterial `json:"low_quality_material_list,omitempty"` //素材评级建议
}

// LowQualityMaterial 低质素材信息
type LowQualityMaterial struct {
	LowQualityVideoIds *[]string `json:"low_quality_video_ids,omitempty"` //低质视频素材
	LowQualityImageIds *[]string `json:"low_quality_image_ids,omitempty"` //低质图片素材
}

// TemplateData 模版自定义参数
type TemplateData struct {
	BackgroundImageId *[]string `json:"background_image_id,omitempty"` //自定义背景图片ID，图片尺寸必须与模版背景图尺寸一致
}

// BrandInfo 品牌信息
type BrandInfo struct {
	YuntuCategoryId *uint64   `json:"yuntu_category_id,omitempty"`  //品牌分类id
	CdpBrandId      *uint64   `json:"cdp_brand_id,omitempty"`       //cdp品牌id
	EcomBrandId     *uint64   `json:"ecom_brand_id,omitempty"`      //电商品牌id
	BrandNameId     *uint64   `json:"brand_name_id,omitempty"`      //云图品牌id
	CdpBrandName    *string   `json:"cdp_brand_name,omitempty"`     //cdp品牌名称
	SubBrandNames   *[]string `json:"sub_brand_names,omitempty"`    //子品牌名称
	SubBrandNameIds *[]string `json:"sub_brand_name_ids,omitempty"` //子品牌id
}

// PromotionKeywords 关键词列表
type PromotionKeywords struct {
	Word      *string   `json:"word,omitempty"` //关键词
	MatchType MatchType `json:"match_type,omitempty"`
	Bid       *string   `json:"bid,omitempty"` //出价
}

// NativeSetting 原生广告设置
type NativeSetting struct {
	AwemeId           *string `json:"aweme_id,omitempty"`
	IsFeedAndFavSee   *Switch `json:"is_feed_and_fav_see,omitempty"` //主页作品列表隐藏广告内容允选值：OFF（不隐藏），ON（隐藏）
	AnchorRelatedType *string `json:"anchor_related_type,omitempty"` // 原生锚点启用开关，允许值: 不启用 OFF，自动生成 AUTO，手动选择 SELECT默认值为 OFF
}
