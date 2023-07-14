package enums

// LandingType : 推广目的类型
type LandingType string

const (
	LikeLandingType     LandingType = "LINK"
	AppLandingType      LandingType = "APP"
	QuickAppLandingType LandingType = "QUICK_APP"
	DpaLandingType      LandingType = "DPA"
	GoodsLandingType    LandingType = "GOODS"
	StoreLandingType    LandingType = "STORE"
	AwemeLandingType    LandingType = "AWEME"
	ShopLandingType     LandingType = "SHOP"
	LiveLandingType     LandingType = "LIVE"
	ArticalLandingType  LandingType = "ARTICAL"
)

func (t LandingType) String() string {
	switch t {
	case LikeLandingType:
		return "销售线索推广"
	case AppLandingType:
		return "应用推广"
	case QuickAppLandingType:
		return "快应用"
	case DpaLandingType:
		return "商品目录"
	case GoodsLandingType:
		return "商品推广（鲁班）"
	case StoreLandingType:
		return "门店推广"
	case AwemeLandingType:
		return "抖音号推广"
	case ShopLandingType:
		return "电商店铺推广"
	case LiveLandingType:
		return "直播间推广"
	case ArticalLandingType:
		return "头条文章推广"
	default:
		return ""
	}
}
