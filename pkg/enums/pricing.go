package enums

// PricingType : 计费方式
type PricingType string

const (
	PricingTypeIsCpm  PricingType = "PRICING_CPM"
	PricingTypeIsCpc  PricingType = "PRICING_CPC"
	PricingTypeIsOcpm PricingType = "PRICING_OCPM"
	PricingTypeIsOcpc PricingType = "PRICING_OCPC"
)

func (t PricingType) String() string {
	switch t {
	case PricingTypeIsCpm:
		return "按展示付费"
	case PricingTypeIsCpc:
		return "按点击付费"
	case PricingTypeIsOcpm:
		return "目标转化出价-按展示付费"
	case PricingTypeIsOcpc:
		return "目标转化出价-按点击付费"
	default:
		return ""
	}
}
