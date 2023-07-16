package enums

// AppPromotionType : 子目标类型
type AppPromotionType string

const (
	AppPromotionTypeIsDOWNLOAD AppPromotionType = "DOWNLOAD"
	AppPromotionTypeIsLAUNCH   AppPromotionType = "LAUNCH"
	AppPromotionTypeIsRESERVE  AppPromotionType = "RESERVE"
)

func (t AppPromotionType) String() string {
	switch t {
	case AppPromotionTypeIsDOWNLOAD:
		return "应用下载"
	case AppPromotionTypeIsLAUNCH:
		return "应用调用"
	case AppPromotionTypeIsRESERVE:
		return "预约下载"
	default:
		return ""
	}
}
