package enums

// Gender : 受众性别
type Gender string

const (
	NoneGender      Gender = "NONE"
	UnlimitedGender Gender = "GENDER_UNLIMITED"
	MaleGender      Gender = "GENDER_MALE"
	FemaleGender    Gender = "GENDER_FEMALE"
)

func (t Gender) String() string {
	switch t {
	case NoneGender:
		return "不限"
	case UnlimitedGender:
		return "不限"
	case MaleGender:
		return "男"
	case FemaleGender:
		return "女"
	default:
		return ""
	}
}
