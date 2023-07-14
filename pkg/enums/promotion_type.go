package enums

// PromotionType : 投放内容
type PromotionType string

const (
	AwemeHomePagePromotionType   PromotionType = "AWEME_HOME_PAGE"
	LandingPageLinkPromotionType PromotionType = "LANDING_PAGE_LINK"
)

func (t PromotionType) String() string {
	switch t {
	case AwemeHomePagePromotionType:
		return "抖音主页"
	case LandingPageLinkPromotionType:
		return "落地页"
	default:
		return ""
	}
}
