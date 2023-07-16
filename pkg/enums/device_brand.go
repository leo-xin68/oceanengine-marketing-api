package enums

// DeviceBrand : 手机品牌
type DeviceBrand string

const (
	DeviceBrandIsHONOR       DeviceBrand = "HONOR"
	DeviceBrandIsAPPLE       DeviceBrand = "APPLE"
	DeviceBrandIsHUAWEI      DeviceBrand = "HUAWEI"
	DeviceBrandIsXIAOMI      DeviceBrand = "XIAOMI"
	DeviceBrandIsSamsung     DeviceBrand = "SAMSUNG"
	DeviceBrandIsOPPO        DeviceBrand = "OPPO"
	DeviceBrandIsVIVO        DeviceBrand = "VIVO"
	DeviceBrandIsMEIZU       DeviceBrand = "MEIZU"
	DeviceBrandIsGIONEE      DeviceBrand = "GIONEE"
	DeviceBrandIsCOOLPAD     DeviceBrand = "COOLPAD"
	DeviceBrandIsLENOVO      DeviceBrand = "LENOVO"
	DeviceBrandIsLETV        DeviceBrand = "LETV"
	DeviceBrandIsZTE         DeviceBrand = "ZTE"
	DeviceBrandIsCHINAMOBILE DeviceBrand = "CHINAMOBILE"
	DeviceBrandIsHTC         DeviceBrand = "HTC"
	DeviceBrandIsPEPPER      DeviceBrand = "PEPPER"
	DeviceBrandIsNUBIA       DeviceBrand = "NUBIA"
	DeviceBrandIsHISENSE     DeviceBrand = "HISENSE"
	DeviceBrandIsQIKU        DeviceBrand = "QIKU"
	DeviceBrandIsTCL         DeviceBrand = "TCL"
	DeviceBrandIsSONY        DeviceBrand = "SONY"
	DeviceBrandIsSMARTISAN   DeviceBrand = "SMARTISAN"
	DeviceBrandIs360         DeviceBrand = "360"
	DeviceBrandIsONEPLUS     DeviceBrand = "ONEPLUS"
	DeviceBrandIsLG          DeviceBrand = "LG"
	DeviceBrandIsMOTO        DeviceBrand = "MOTO"
	DeviceBrandIsNOKIA       DeviceBrand = "NOKIA"
	DeviceBrandIsGOOGLE      DeviceBrand = "GOOGLE"
)

func (t DeviceBrand) String() string {
	switch t {
	case DeviceBrandIsHONOR:
		return "荣耀"
	case DeviceBrandIsAPPLE:
		return "苹果"
	case DeviceBrandIsHUAWEI:
		return "华为"
	case DeviceBrandIsXIAOMI:
		return "小米"
	case DeviceBrandIsSamsung:
		return "三星"
	case DeviceBrandIsOPPO:
		return "OPPO"
	case DeviceBrandIsVIVO:
		return "VIVO"
	case DeviceBrandIsMEIZU:
		return "魅族"
	case DeviceBrandIsGIONEE:
		return "金立"
	case DeviceBrandIsCOOLPAD:
		return "酷派"
	case DeviceBrandIsLENOVO:
		return "联想"
	case DeviceBrandIsLETV:
		return "乐视"
	case DeviceBrandIsZTE:
		return "中兴"
	case DeviceBrandIsCHINAMOBILE:
		return "中国移动"
	case DeviceBrandIsHTC:
		return "HTC"
	case DeviceBrandIsPEPPER:
		return "小辣椒"
	case DeviceBrandIsNUBIA:
		return "努比亚"
	case DeviceBrandIsHISENSE:
		return "海信"
	case DeviceBrandIsQIKU:
		return "奇酷"
	case DeviceBrandIsTCL:
		return "TCL"
	case DeviceBrandIsSONY:
		return "索尼"
	case DeviceBrandIsSMARTISAN:
		return "锤子手机"
	case DeviceBrandIs360:
		return "360手机"
	case DeviceBrandIsONEPLUS:
		return "一加手机"
	case DeviceBrandIsLG:
		return "LG"
	case DeviceBrandIsMOTO:
		return "摩托罗拉"
	case DeviceBrandIsNOKIA:
		return "诺基亚"
	case DeviceBrandIsGOOGLE:
		return "谷歌"
	default:
		return ""
	}
}
