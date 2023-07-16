package enums

// DeepBidType : 深度优化方式
type DeepBidType string

const (
	DeepBidTypeIsDeepBidMin     DeepBidType = "DEEP_BID_MIN"
	DeepBidTypeIsRoiCoefficient DeepBidType = "ROI_COEFFICIENT"
	DeepBidTypeIsBidPerAction   DeepBidType = "BID_PER_ACTION"
	DeepBidTypeIsSocialRoi      DeepBidType = "SOCIAL_ROI"
)

func (t DeepBidType) String() string {
	switch t {
	case DeepBidTypeIsDeepBidMin:
		return "自定义手动出价"
	case DeepBidTypeIsRoiCoefficient:
		return "ROI系数出价"
	case DeepBidTypeIsBidPerAction:
		return "每次付费出价"
	case DeepBidTypeIsSocialRoi:
		return "ROI三出价"
	default:
		return ""
	}
}
