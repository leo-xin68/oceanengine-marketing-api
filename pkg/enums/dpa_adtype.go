package enums

// DpaAdType : DPA广告类型
type DpaAdType string

const (
	DpaLinkDpaAdType DpaAdType = "DPA_LINK"
	DpaAppDpaAdType  DpaAdType = "DPA_APP"
)

func (t DpaAdType) String() string {
	switch t {
	case DpaLinkDpaAdType:
		return "落地页"
	case DpaAppDpaAdType:
		return "应用下载"
	default:
		return ""
	}
}
