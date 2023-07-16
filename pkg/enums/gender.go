package enums

// Gender : 受众性别
type Gender string

const (
	GenderIsNone      Gender = "NONE"
	GenderIsUnlimited Gender = "GENDER_UNLIMITED"
	GenderIsMale      Gender = "GENDER_MALE"
	GenderIsFemale    Gender = "GENDER_FEMALE"
)

func (t Gender) String() string {
	switch t {
	case GenderIsNone:
		return "不限"
	case GenderIsUnlimited:
		return "不限"
	case GenderIsMale:
		return "男"
	case GenderIsFemale:
		return "女"
	default:
		return ""
	}
}
