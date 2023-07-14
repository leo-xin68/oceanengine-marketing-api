package enums

// AcType : 受众网络类型
type AcType string

const (
	AcTypeUnknown AcType = "unknown"
	AcTypeWifi    AcType = "WIFI"
	AcType2G      AcType = "2G"
	AcType3G      AcType = "3G"
	AcType4G      AcType = "4G"
	AcType5G      AcType = "5G"
)

func (t AcType) String() string {
	switch t {
	case AcTypeUnknown:
		return "unknown"
	case AcTypeWifi:
		return "wifi"
	case AcType2G:
		return "2G"
	case AcType3G:
		return "3G"
	case AcType4G:
		return "4G"
	case AcType5G:
		return "5G"
	default:
		return ""
	}
}
