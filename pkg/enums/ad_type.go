package enums

// AdType : 广告类型
type AdType string

const (
	AdTypeIsAll    AdType = "ALL"
	AdTypeIsSearch AdType = "SEARCH"
)

func (t AdType) String() string {
	switch t {
	case AdTypeIsAll:
		return "通投广告"
	case AdTypeIsSearch:
		return "搜索广告"
	default:
		return ""
	}
}

func (t AdType) Check() bool {
	switch t {
	case AdTypeIsAll:
		return true
	case AdTypeIsSearch:
		return true
	default:
		return false
	}
}
