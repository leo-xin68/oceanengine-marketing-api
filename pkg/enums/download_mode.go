package enums

// DownloadMode : 下载模式
type DownloadMode string

const (
	AppStoreDeliveryDownloadMode DownloadMode = "APP_STORE_DELIVERY"
	DefaultDownloadMode          DownloadMode = "DEFAULT"
)

func (t DownloadMode) String() string {
	switch t {
	case AppStoreDeliveryDownloadMode:
		return "商店下载"
	case DefaultDownloadMode:
		return "默认下载"
	default:
		return ""
	}
}
