package enums

// LandingType : 推广目的类型
type LandingType string

const (
	LandingTypeIsLike     LandingType = "LINK"
	LandingTypeIsApp      LandingType = "APP"
	LandingTypeIsQuickApp LandingType = "QUICK_APP"
	LandingTypeIsDpa      LandingType = "DPA"
	LandingTypeIsGoods    LandingType = "GOODS"
	LandingTypeIsStore    LandingType = "STORE"
	LandingTypeIsAweme    LandingType = "AWEME"
	LandingTypeIsShop     LandingType = "SHOP"
	LandingTypeIsLive     LandingType = "LIVE"
	LandingTypeIsArtical  LandingType = "ARTICAL"
)

func (t LandingType) String() string {
	switch t {
	case LandingTypeIsLike:
		return "销售线索推广"
	case LandingTypeIsApp:
		return "应用推广"
	case LandingTypeIsQuickApp:
		return "快应用"
	case LandingTypeIsDpa:
		return "商品目录"
	case LandingTypeIsGoods:
		return "商品推广（鲁班）"
	case LandingTypeIsStore:
		return "门店推广"
	case LandingTypeIsAweme:
		return "抖音号推广"
	case LandingTypeIsShop:
		return "电商店铺推广"
	case LandingTypeIsLive:
		return "直播间推广"
	case LandingTypeIsArtical:
		return "头条文章推广"
	default:
		return ""
	}
}
