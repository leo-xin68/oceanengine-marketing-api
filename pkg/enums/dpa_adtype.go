package enums

// DpaAdType : DPA广告类型
type DpaAdType string

const (
	DpaAdTypeIsDpaLink DpaAdType = "DPA_LINK"
	DpaAdTypeIsDpaApp  DpaAdType = "DPA_APP"
)

func (t DpaAdType) String() string {
	switch t {
	case DpaAdTypeIsDpaLink:
		return "落地页"
	case DpaAdTypeIsDpaApp:
		return "应用下载"
	default:
		return ""
	}
}
