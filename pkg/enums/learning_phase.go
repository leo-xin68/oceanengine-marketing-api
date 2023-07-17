package enums

// LearningPhase : 学习期状态
type LearningPhase string

const (
	LearningPhaseIsDefault     LearningPhase = "DEFAULT"
	LearningPhaseIsLearning    LearningPhase = "LEARNING"
	LearningPhaseIsLearned     LearningPhase = "LEARNED"
	LearningPhaseIsLearnFailed LearningPhase = "LEARN_FAILED"
)

func (t LearningPhase) String() string {
	switch t {
	case LearningPhaseIsDefault:
		return "默认，不在学习期中"
	case LearningPhaseIsLearning:
		return "学习中"
	case LearningPhaseIsLearned:
		return "学习成功"
	case LearningPhaseIsLearnFailed:
		return "学习失败"
	default:
		return ""
	}
}
