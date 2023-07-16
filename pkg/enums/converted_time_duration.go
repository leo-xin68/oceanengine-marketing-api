package enums

// ConvertedTimeDuration : 过滤时间范围
type ConvertedTimeDuration string

const (
	ConvertedTimeDurationIsNone        ConvertedTimeDuration = "NONE"
	ConvertedTimeDurationIsToday       ConvertedTimeDuration = "TODAY"
	ConvertedTimeDurationIsSevenDay    ConvertedTimeDuration = "SEVEN_DAY"
	ConvertedTimeDurationIsOneMonth    ConvertedTimeDuration = "ONE_MONTH"
	ConvertedTimeDurationIsThreeMonth  ConvertedTimeDuration = "THREE_MONTH"
	ConvertedTimeDurationIsSixMonth    ConvertedTimeDuration = "SIX_MONTH"
	ConvertedTimeDurationIsTwelveMonth ConvertedTimeDuration = "TWELVE_MONTH"
)

func (t ConvertedTimeDuration) String() string {
	switch t {
	case ConvertedTimeDurationIsNone:
		return "不限"
	case ConvertedTimeDurationIsToday:
		return "当天"
	case ConvertedTimeDurationIsSevenDay:
		return "7天"
	case ConvertedTimeDurationIsOneMonth:
		return "1个月"
	case ConvertedTimeDurationIsThreeMonth:
		return "3个月"
	case ConvertedTimeDurationIsSixMonth:
		return "6个月"
	case ConvertedTimeDurationIsTwelveMonth:
		return "12个月"
	default:
		return ""
	}
}
