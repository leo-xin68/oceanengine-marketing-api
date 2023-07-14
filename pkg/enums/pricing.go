package enums

// PricingType : 计费方式
type PricingType string

const (
	CpmPricingType  PricingType = "PRICING_CPM"
	CpcPricingType  PricingType = "PRICING_CPC"
	OcpmPricingType PricingType = "PRICING_OCPM"
	OcpcPricingType PricingType = "PRICING_OCPC"
)

func (t PricingType) String() string {
	switch t {
	case CpmPricingType:
		return "按展示付费"
	case CpcPricingType:
		return "按点击付费"
	case OcpmPricingType:
		return "目标转化出价-按展示付费"
	case OcpcPricingType:
		return "目标转化出价-按点击付费"
	default:
		return ""
	}
}
