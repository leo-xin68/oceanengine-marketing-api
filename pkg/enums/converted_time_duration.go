package enums

// ConvertedTimeDuration : 过滤时间范围
type ConvertedTimeDuration string

const (
	ConvertedTimeDuration_NONE         ConvertedTimeDuration = "NONE"
	ConvertedTimeDuration_TODAY        ConvertedTimeDuration = "TODAY"
	ConvertedTimeDuration_SEVEN_DAY    ConvertedTimeDuration = "SEVEN_DAY"
	ConvertedTimeDuration_ONE_MONTH    ConvertedTimeDuration = "ONE_MONTH"
	ConvertedTimeDuration_THREE_MONTH  ConvertedTimeDuration = "THREE_MONTH"
	ConvertedTimeDuration_SIX_MONTH    ConvertedTimeDuration = "SIX_MONTH"
	ConvertedTimeDuration_TWELVE_MONTH ConvertedTimeDuration = "TWELVE_MONTH"
)
