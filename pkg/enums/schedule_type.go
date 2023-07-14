package enums

// ScheduleType : 投放时间类型
type ScheduleType string

const (
	FromNowScheduleType  ScheduleType = "SCHEDULE_FROM_NOW"
	StartEndScheduleType ScheduleType = "SCHEDULE_START_END"
)

func (t ScheduleType) String() string {
	switch t {
	case FromNowScheduleType:
		return "从今天起长期投放"
	case StartEndScheduleType:
		return "设置开始和结束日期"
	default:
		return ""
	}
}
