package enums

// TrackUrlType : 监测链接类型
type TrackUrlType string

const (
	CustomTrackUrlType  TrackUrlType = "CUSTOM"
	GroupIdTrackUrlType TrackUrlType = "GROUP_ID"
)

func (t TrackUrlType) String() string {
	switch t {
	case CustomTrackUrlType:
		return "自定义链接"
	case GroupIdTrackUrlType:
		return "监测链接组"
	default:
		return ""
	}
}
