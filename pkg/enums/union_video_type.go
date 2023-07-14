package enums

// UnionVideoType : 穿山甲视频创意类型
type UnionVideoType string

const (
	OriginalVideoUnionVideoType UnionVideoType = "ORIGINAL_VIDEO"
	RewardedVideoUnionVideoType UnionVideoType = "REWARDED_VIDEO"
	SplashVideoUnionVideoType   UnionVideoType = "SPLASH_VIDEO"
)

func (t UnionVideoType) String() string {
	switch t {
	case OriginalVideoUnionVideoType:
		return "原生视频"
	case RewardedVideoUnionVideoType:
		return "激励视频"
	case SplashVideoUnionVideoType:
		return "开屏视频"
	default:
		return ""
	}
}
