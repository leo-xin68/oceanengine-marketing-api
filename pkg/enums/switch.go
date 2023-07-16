package enums

// Switch : 开关
type Switch string

const (
	SwitchIsOn  Switch = "ON"
	SwitchIsOff Switch = "OFF"
)

func (t Switch) String() string {
	switch t {
	case SwitchIsOn:
		return "启用"
	case SwitchIsOff:
		return "不启用"
	default:
		return ""
	}
}
