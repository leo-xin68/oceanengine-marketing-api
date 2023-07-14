package enums

// OptimizedType : 目标优化类型
type OptimizedType string

const (
	OffOptimizedType    OptimizedType = "OFF"
	ActionOptimizedType OptimizedType = "ACTION"
	ValueOptimizedType  OptimizedType = "VALUE"
)

func (t OptimizedType) String() string {
	switch t {
	case OffOptimizedType:
		return "不启用"
	case ActionOptimizedType:
		return "行为优化"
	case ValueOptimizedType:
		return "价值优化"
	default:
		return ""
	}
}
