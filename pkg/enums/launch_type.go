package enums

// LaunchType : 调起方式
type LaunchType string

const (
	DirectOpenLaunch   LaunchType = "DIRECT_OPEN"
	ExternalOpenLaunch LaunchType = "EXTERNAL_OPEN"
)

func (t LaunchType) String() string {
	switch t {
	case DirectOpenLaunch:
		return "直接调起"
	case ExternalOpenLaunch:
		return "落地页调起"
	default:
		return ""
	}
}
