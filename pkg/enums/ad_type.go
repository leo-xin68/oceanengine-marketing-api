package enums

// AdType : 广告类型
type AdType string

const (
	AllAd    AdType = "ALL"
	SearchAd AdType = "SEARCH"
)

func (t AdType) String() string {
	switch t {
	case AllAd:
		return "通投广告"
	case SearchAd:
		return "搜索广告"
	default:
		return ""
	}
}
