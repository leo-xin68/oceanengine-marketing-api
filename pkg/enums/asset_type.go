package enums

// AssetType : 资产类型
type AssetType string

const (
	AssetTypeIsOrange        AssetType = "ORANGE"
	AssetTypeIsThirdparty    AssetType = "THIRDPARTY"
	AssetTypeIsEnterprise    AssetType = "ENTERPRISE"
	AssetTypeIsAwemeHomePage AssetType = "AWEME_HOME_PAGE"
)

func (t AssetType) String() string {
	switch t {
	case AssetTypeIsOrange:
		return "橙子落地页"
	case AssetTypeIsThirdparty:
		return "自研落地页"
	case AssetTypeIsEnterprise:
		return "企业号落地页"
	case AssetTypeIsAwemeHomePage:
		return "抖音主页"
	default:
		return ""
	}
}
