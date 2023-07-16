package enums

// MicroPromotionType : 小程序类型
type MicroPromotionType string

const (
	MicroPromotionTypeIsWechatGame MicroPromotionType = "WECHAT_GAME"
	MicroPromotionTypeIsWechatApp  MicroPromotionType = "WECHAT_APP"
	MicroPromotionTypeIsByteGame   MicroPromotionType = "BYTE_GAME"
	MicroPromotionTypeIsByteApp    MicroPromotionType = "BYTE_APP"
)

func (t MicroPromotionType) String() string {
	switch t {
	case MicroPromotionTypeIsWechatGame:
		return "微信小游戏"
	case MicroPromotionTypeIsWechatApp:
		return "微信小程序"
	case MicroPromotionTypeIsByteGame:
		return "字节小游戏"
	case MicroPromotionTypeIsByteApp:
		return "字节小程序"
	default:
		return ""
	}
}
