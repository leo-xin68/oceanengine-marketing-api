package enums

// MicroPromotionType : 小程序类型
type MicroPromotionType string

const (
	WechatGameMicroPromotionType MicroPromotionType = "WECHAT_GAME"
	WechatAppMicroPromotionType  MicroPromotionType = "WECHAT_APP"
	ByteGameMicroPromotionType   MicroPromotionType = "BYTE_GAME"
	ByteAppMicroPromotionType    MicroPromotionType = "BYTE_APP"
)

func (t MicroPromotionType) String() string {
	switch t {
	case WechatGameMicroPromotionType:
		return "微信小游戏"
	case WechatAppMicroPromotionType:
		return "微信小程序"
	case ByteGameMicroPromotionType:
		return "字节小游戏"
	case ByteAppMicroPromotionType:
		return "字节小程序"
	default:
		return ""
	}
}
