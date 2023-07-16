package enums

// UnionVideoType : 穿山甲视频创意类型
type UnionVideoType string

const (
	UnionVideoTypeIsOriginalVideo UnionVideoType = "ORIGINAL_VIDEO"
	UnionVideoTypeIsRewardedVideo UnionVideoType = "REWARDED_VIDEO"
	UnionVideoTypeIsSplashVideo   UnionVideoType = "SPLASH_VIDEO"
)

func (t UnionVideoType) String() string {
	switch t {
	case UnionVideoTypeIsOriginalVideo:
		return "原生视频"
	case UnionVideoTypeIsRewardedVideo:
		return "激励视频"
	case UnionVideoTypeIsSplashVideo:
		return "开屏视频"
	default:
		return ""
	}
}
