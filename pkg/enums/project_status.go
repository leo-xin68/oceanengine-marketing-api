package enums

// ProjectStatus : 广告项目状态
type ProjectStatus string

const (
	ProjectStatusIsEnable                 ProjectStatus = "PROJECT_STATUS_ENABLE"
	ProjectStatusIsDisable                ProjectStatus = "PROJECT_STATUS_DISABLE"
	ProjectStatusIsDelete                 ProjectStatus = "PROJECT_STATUS_DELETE"
	ProjectStatusIsAll                    ProjectStatus = "PROJECT_STATUS_ALL"
	ProjectStatusIsNotDelete              ProjectStatus = "PROJECT_STATUS_NOT_DELETE"
	ProjectStatusIsBudgetExceed           ProjectStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	ProjectStatusIsBudgetPreOfflineBudget ProjectStatus = "PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET"
	ProjectStatusIsNotStart               ProjectStatus = "PROJECT_STATUS_NOT_START"
	ProjectStatusIsDone                   ProjectStatus = "PROJECT_STATUS_DONE"
	ProjectStatusIsNoSchedule             ProjectStatus = "PROJECT_STATUS_NO_SCHEDULE"
)

func (t ProjectStatus) String() string {
	switch t {
	case ProjectStatusIsEnable:
		return "启用"
	case ProjectStatusIsDisable:
		return "暂停"
	case ProjectStatusIsDelete:
		return "删除"
	case ProjectStatusIsAll:
		return "所有包含已删除"
	case ProjectStatusIsNotDelete:
		return "所有不包含已删除"
	case ProjectStatusIsBudgetExceed:
		return "项目超出预算"
	case ProjectStatusIsBudgetPreOfflineBudget:
		return "项目接近预算"
	case ProjectStatusIsNotStart:
		return "未达投放时间"
	case ProjectStatusIsDone:
		return "已完成"
	case ProjectStatusIsNoSchedule:
		return "不在投放时段"
	default:
		return ""
	}
}

// ProjectFirstStatus 项目一级状态
type ProjectFirstStatus string

const (
	ProjectFirstStatusIsDelete  ProjectFirstStatus = "PROJECT_STATUS_DELETE"
	ProjectFirstStatusIsDone    ProjectFirstStatus = "PROJECT_STATUS_DONE"
	ProjectFirstStatusIsDisable ProjectFirstStatus = "PROJECT_STATUS_DISABLE"
	ProjectFirstStatusIsEnable  ProjectFirstStatus = "PROJECT_STATUS_ENABLE"
)

func (t ProjectFirstStatus) String() string {
	switch t {
	case ProjectFirstStatusIsDelete:
		return "已删除"
	case ProjectFirstStatusIsDone:
		return "已完成"
	case ProjectFirstStatusIsDisable:
		return "未投放"
	case ProjectFirstStatusIsEnable:
		return "启用中"
	default:
		return ""
	}
}

// ProjectSecondStatus 项目二级状态
type ProjectSecondStatus string

const (
	ProjectSecondStatusIsStop         ProjectSecondStatus = "PROJECT_STATUS_STOP"
	ProjectSecondStatusIsBudgetExceed ProjectSecondStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	ProjectSecondStatusIsNotStart     ProjectSecondStatus = "PROJECT_STATUS_NOT_START"
	ProjectSecondStatusIsNoSchedule   ProjectSecondStatus = "PROJECT_STATUS_NO_SCHEDULE"
)

func (t ProjectSecondStatus) String() string {
	switch t {
	case ProjectSecondStatusIsStop:
		return "已暂停"
	case ProjectSecondStatusIsBudgetExceed:
		return "项目超出预算"
	case ProjectSecondStatusIsNotStart:
		return "未达投放时间"
	case ProjectSecondStatusIsNoSchedule:
		return "不在投放时段"
	default:
		return ""
	}
}
