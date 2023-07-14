package enums

// AssetType : 资产类型
type AssetType string

const (
	OrangeAssetType        AssetType = "ORANGE"
	ThirdpartyAssetType    AssetType = "THIRDPARTY"
	EnterpriseAssetType    AssetType = "ENTERPRISE"
	AwemeHomePageAssetType AssetType = "AWEME_HOME_PAGE"
)

func (t AssetType) String() string {
	switch t {
	case OrangeAssetType:
		return "橙子落地页"
	case ThirdpartyAssetType:
		return "自研落地页"
	case EnterpriseAssetType:
		return "企业号落地页"
	case AwemeHomePageAssetType:
		return "抖音主页"
	default:
		return ""
	}
}
