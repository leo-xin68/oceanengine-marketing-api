package enums

// MatchType : 广告类型
type MatchType string

const (
	MatchTypeIsPhrase    MatchType = "PHRASE"
	MatchTypeIsExtensive MatchType = "EXTENSIVE"
	MatchTypeIsPrecision MatchType = "PRECISION"
)

func (t MatchType) String() string {
	switch t {
	case MatchTypeIsPhrase:
		return "短语匹配"
	case MatchTypeIsExtensive:
		return "广泛匹配"
	case MatchTypeIsPrecision:
		return "精准匹配"
	default:
		return ""
	}
}
