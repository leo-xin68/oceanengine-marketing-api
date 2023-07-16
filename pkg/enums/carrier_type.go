package enums

// CarrierType : 受众运营商类型
type CarrierType string

const (
	CarrierTypeIsMobile CarrierType = "MOBILE"
	CarrierTypeIsUnicom CarrierType = "UNICOM"
	CarrierTypeIsTelcom CarrierType = "TELCOM"
)

func (t CarrierType) String() string {
	switch t {
	case CarrierTypeIsMobile:
		return "移动"
	case CarrierTypeIsUnicom:
		return "联通"
	case CarrierTypeIsTelcom:
		return "电信"
	default:
		return ""
	}
}
