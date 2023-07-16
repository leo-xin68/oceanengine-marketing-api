package enums

// PromotionType : 投放内容
type PromotionType string

const (
	PromotionTypeIsAwemeHomePage   PromotionType = "AWEME_HOME_PAGE"
	PromotionTypeIsLandingPageLink PromotionType = "LANDING_PAGE_LINK"
)

func (t PromotionType) String() string {
	switch t {
	case PromotionTypeIsAwemeHomePage:
		return "抖音主页"
	case PromotionTypeIsLandingPageLink:
		return "落地页"
	default:
		return ""
	}
}
