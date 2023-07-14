package enums

// AgeInterval : 受众年龄区间
type AgeInterval string

const (
	Between18To23AgeInterval AgeInterval = "AGE_BETWEEN_18_23"
	Between24To30AgeInterval AgeInterval = "AGE_BETWEEN_24_30"
	Between31To40AgeInterval AgeInterval = "AGE_BETWEEN_31_40"
	Between41To49AgeInterval AgeInterval = "AGE_BETWEEN_41_49"
	Above50AgeInterval       AgeInterval = "AGE_ABOVE_50"
)

func (t AgeInterval) String() string {
	switch t {
	case Between18To23AgeInterval:
		return "18-23岁"
	case Between24To30AgeInterval:
		return "24-30岁"
	case Between31To40AgeInterval:
		return "31-40岁"
	case Between41To49AgeInterval:
		return "41-49岁"
	case Above50AgeInterval:
		return "大于等于50岁"
	default:
		return ""
	}
}
