package enums

// MaterialsType : 素材类型
type MaterialsType string

const (
	MaterialsTypeIsLiveMaterials      MaterialsType = "LIVE_MATERIALS"
	MaterialsTypeIsPromotionMaterials MaterialsType = "PROMOTION_MATERIALS"
)

func (t MaterialsType) String() string {
	switch t {
	case MaterialsTypeIsLiveMaterials:
		return "直播素材"
	case MaterialsTypeIsPromotionMaterials:
		return "广告素材"
	default:
		return ""
	}
}
