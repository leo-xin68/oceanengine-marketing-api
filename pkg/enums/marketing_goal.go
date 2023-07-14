package enums

// MarketingGoal : 营销场景
type MarketingGoal string

const (
	VideoAndImageMarketingGoal MarketingGoal = "VIDEO_AND_IMAGE"
	LiveMarketingGoal          MarketingGoal = "LIVE"
)

func (t MarketingGoal) String() string {
	switch t {
	case VideoAndImageMarketingGoal:
		return "短视频/图片"
	case LiveMarketingGoal:
		return "直播"
	default:
		return ""
	}
}
