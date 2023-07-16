package enums

// OptimizedType : 目标优化类型
type OptimizedType string

const (
	OptimizedTypeIsOff    OptimizedType = "OFF"
	OptimizedTypeIsAction OptimizedType = "ACTION"
	OptimizedTypeIsValue  OptimizedType = "VALUE"
)

func (t OptimizedType) String() string {
	switch t {
	case OptimizedTypeIsOff:
		return "不启用"
	case OptimizedTypeIsAction:
		return "行为优化"
	case OptimizedTypeIsValue:
		return "价值优化"
	default:
		return ""
	}
}
