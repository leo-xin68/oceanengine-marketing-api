package enums

// DownloadMode : 下载模式
type DownloadMode string

const (
	DownloadModeIsAppStoreDelivery DownloadMode = "APP_STORE_DELIVERY"
	DownloadModeIsDefault          DownloadMode = "DEFAULT"
)

func (t DownloadMode) String() string {
	switch t {
	case DownloadModeIsAppStoreDelivery:
		return "商店下载"
	case DownloadModeIsDefault:
		return "默认下载"
	default:
		return ""
	}
}
