package enums

// ProjectStatus : 广告项目状态
type ProjectStatus string

const (
	ProjectStatus_ENABLE                    ProjectStatus = "PROJECT_STATUS_ENABLE"
	ProjectStatus_DISABLE                   ProjectStatus = "PROJECT_STATUS_DISABLE"
	ProjectStatus_DELETE                    ProjectStatus = "PROJECT_STATUS_DELETE"
	ProjectStatus_ALL                       ProjectStatus = "PROJECT_STATUS_ALL"
	ProjectStatus_NOT_DELETE                ProjectStatus = "PROJECT_STATUS_NOT_DELETE"
	ProjectStatus_BUDGET_EXCEED             ProjectStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	ProjectStatus_BUDGET_PRE_OFFLINE_BUDGET ProjectStatus = "PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET"
	ProjectStatus_NOT_START                 ProjectStatus = "PROJECT_STATUS_NOT_START"
	ProjectStatus_DONE                      ProjectStatus = "PROJECT_STATUS_DONE"
	ProjectStatus_NO_SCHEDULE               ProjectStatus = "PROJECT_STATUS_NO_SCHEDULE"
)

// ProjectFirstStatus 项目一级状态
type ProjectFirstStatus string

const (
	ProjectFirstStatus_DELETE  ProjectStatus = "PROJECT_STATUS_DELETE"
	ProjectFirstStatus_DONE    ProjectStatus = "PROJECT_STATUS_DONE"
	ProjectFirstStatus_DISABLE ProjectStatus = "PROJECT_STATUS_DISABLE"
	ProjectFirstStatus_ENABLE  ProjectStatus = "PROJECT_STATUS_ENABLE"
)

// ProjectSecondStatus 项目二级状态
type ProjectSecondStatus string

const (
	ProjectSecondStatus_STOP          ProjectStatus = "PROJECT_STATUS_STOP"
	ProjectSecondStatus_BUDGET_EXCEED ProjectStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	ProjectSecondStatus_NOT_START     ProjectStatus = "PROJECT_STATUS_NOT_START"
	ProjectSecondStatus_NO_SCHEDULE   ProjectStatus = "PROJECT_STATUS_NO_SCHEDULE"
)
