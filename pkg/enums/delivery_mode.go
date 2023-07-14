package enums

// DeliveryMode : 投放模式
type DeliveryMode string

const (
	ManualDeliveryMode     DeliveryMode = "MANUAL"
	ProceduralDeliveryMode DeliveryMode = "PROCEDURAL"
)

func (t DeliveryMode) String() string {
	switch t {
	case ManualDeliveryMode:
		return "手动投放"
	case ProceduralDeliveryMode:
		return "自动投放"
	default:
		return ""
	}
}
