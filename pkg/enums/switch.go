package enums

// Switch : 开关
type Switch string

const (
	SwitchOn  Switch = "ON"
	SwitchOff Switch = "OFF"
)

func (t Switch) String() string {
	switch t {
	case SwitchOn:
		return "启用"
	case SwitchOff:
		return "不启用"
	default:
		return ""
	}
}
