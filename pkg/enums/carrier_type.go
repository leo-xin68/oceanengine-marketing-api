package enums

// CarrierType : 受众运营商类型
type CarrierType string

const (
	MobileCarrierType CarrierType = "MOBILE"
	UnicomCarrierType CarrierType = "UNICOM"
	TelcomCarrierType CarrierType = "TELCOM"
)

func (t CarrierType) String() string {
	switch t {
	case MobileCarrierType:
		return "移动"
	case UnicomCarrierType:
		return "联通"
	case TelcomCarrierType:
		return "电信"
	default:
		return ""
	}
}
