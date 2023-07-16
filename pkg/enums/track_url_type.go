package enums

// TrackUrlType : 监测链接类型
type TrackUrlType string

const (
	TrackUrlTypeIsCustom  TrackUrlType = "CUSTOM"
	TrackUrlTypeIsGroupId TrackUrlType = "GROUP_ID"
)

func (t TrackUrlType) String() string {
	switch t {
	case TrackUrlTypeIsCustom:
		return "自定义链接"
	case TrackUrlTypeIsGroupId:
		return "监测链接组"
	default:
		return ""
	}
}
