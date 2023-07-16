package enums

// AcType : 受众网络类型
type AcType string

const (
	AcTypeIsUnknown AcType = "unknown"
	AcTypeIsWifi    AcType = "WIFI"
	AcTypeIs2G      AcType = "2G"
	AcTypeIs3G      AcType = "3G"
	AcTypeIs4G      AcType = "4G"
	AcTypeIs5G      AcType = "5G"
)

func (t AcType) String() string {
	switch t {
	case AcTypeIsUnknown:
		return "unknown"
	case AcTypeIsWifi:
		return "wifi"
	case AcTypeIs2G:
		return "2G"
	case AcTypeIs3G:
		return "3G"
	case AcTypeIs4G:
		return "4G"
	case AcTypeIs5G:
		return "5G"
	default:
		return ""
	}
}
