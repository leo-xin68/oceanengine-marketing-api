package enums

// AppPromotionType : 子目标类型
type AppPromotionType string

const (
	DownloadAppPromotionType AppPromotionType = "DOWNLOAD"
	LaunchAppPromotionType   AppPromotionType = "LAUNCH"
	ReserveAppPromotionType  AppPromotionType = "RESERVE"
)

func (t AppPromotionType) String() string {
	switch t {
	case DownloadAppPromotionType:
		return "应用下载"
	case LaunchAppPromotionType:
		return "应用调用"
	case ReserveAppPromotionType:
		return "预约下载"
	default:
		return ""
	}
}
