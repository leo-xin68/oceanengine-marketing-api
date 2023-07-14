package enums

// DeepBidType : 深度优化方式
type DeepBidType string

const (
	DeepBidMinDeepBidType     DeepBidType = "DEEP_BID_MIN"
	RoiCoefficientDeepBidType DeepBidType = "ROI_COEFFICIENT"
	BidPerActionDeepBidType   DeepBidType = "BID_PER_ACTION"
	SocialRoiDeepBidType      DeepBidType = "SOCIAL_ROI"
)

func (t DeepBidType) String() string {
	switch t {
	case DeepBidMinDeepBidType:
		return "自定义手动出价"
	case RoiCoefficientDeepBidType:
		return "ROI系数出价"
	case BidPerActionDeepBidType:
		return "每次付费出价"
	case SocialRoiDeepBidType:
		return "ROI三出价"
	default:
		return ""
	}
}
