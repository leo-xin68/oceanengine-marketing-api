package enums

// AgeInterval : 受众年龄区间
type AgeInterval string

const (
	AgeIntervalIsBetween18To23 AgeInterval = "AGE_BETWEEN_18_23"
	AgeIntervalIsBetween24To30 AgeInterval = "AGE_BETWEEN_24_30"
	AgeIntervalIsBetween31To40 AgeInterval = "AGE_BETWEEN_31_40"
	AgeIntervalIsBetween41To49 AgeInterval = "AGE_BETWEEN_41_49"
	AgeIntervalIsAbove50       AgeInterval = "AGE_ABOVE_50"
)

func (t AgeInterval) String() string {
	switch t {
	case AgeIntervalIsBetween18To23:
		return "18-23岁"
	case AgeIntervalIsBetween24To30:
		return "24-30岁"
	case AgeIntervalIsBetween31To40:
		return "31-40岁"
	case AgeIntervalIsBetween41To49:
		return "41-49岁"
	case AgeIntervalIsAbove50:
		return "大于等于50岁"
	default:
		return ""
	}
}
