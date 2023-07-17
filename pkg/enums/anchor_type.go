package enums

// AnchorType : 锚点类型
type AnchorType string

const (
	AnchorTypeIsAppGame            AnchorType = "APP_GAME"
	AnchorTypeIsAppInternetService AnchorType = "APP_INTERNET_SERVICE"
	AnchorTypeIsAppShop            AnchorType = "APP_SHOP"
	AnchorTypeIsOnlineSubscribe    AnchorType = "ONLINE_SUBSCRIBE"
	AnchorTypeIsShoppingCart       AnchorType = "SHOPPING_CART"
	AnchorTypeIsPrivateChat        AnchorType = "PRIVATE_CHAT"
	AnchorTypeIsInsurance          AnchorType = "INSURANCE"
)

func (t AnchorType) String() string {
	switch t {
	case AnchorTypeIsAppGame:
		return "应用下载-游戏"
	case AnchorTypeIsAppInternetService:
		return "应用下载-网服"
	case AnchorTypeIsAppShop:
		return "应用下载-电商"
	case AnchorTypeIsOnlineSubscribe:
		return "高级在线预约"
	case AnchorTypeIsShoppingCart:
		return "购物车"
	case AnchorTypeIsPrivateChat:
		return "私信"
	case AnchorTypeIsInsurance:
		return "保险"
	default:
		return ""
	}
}
