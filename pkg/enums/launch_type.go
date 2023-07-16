package enums

// LaunchType : 调起方式
type LaunchType string

const (
	LaunchTypeIsDirectOpen   LaunchType = "DIRECT_OPEN"
	LaunchTypeIsExternalOpen LaunchType = "EXTERNAL_OPEN"
)

func (t LaunchType) String() string {
	switch t {
	case LaunchTypeIsDirectOpen:
		return "直接调起"
	case LaunchTypeIsExternalOpen:
		return "落地页调起"
	default:
		return ""
	}
}
