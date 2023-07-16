package enums

// MarketingGoal : 营销场景
type MarketingGoal string

const (
	MarketingGoalIsVideoAndImage MarketingGoal = "VIDEO_AND_IMAGE"
	MarketingGoalIsLive          MarketingGoal = "LIVE"
)

func (t MarketingGoal) String() string {
	switch t {
	case MarketingGoalIsVideoAndImage:
		return "短视频/图片"
	case MarketingGoalIsLive:
		return "直播"
	default:
		return ""
	}
}
