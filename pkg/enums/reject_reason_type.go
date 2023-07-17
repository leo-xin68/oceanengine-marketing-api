package enums

// RejectReasonType : 审核建议类型
type RejectReasonType string

const (
	RejectReasonTypeIsNone         RejectReasonType = "NONE"
	RejectReasonTypeIsReview       RejectReasonType = "REVIEW_REJECT"
	RejectReasonTypeIsLowMaterail  RejectReasonType = "LOW_MATERAIL"
	RejectReasonTypeIsDiscomfort   RejectReasonType = "DISCOMFORT"
	RejectReasonTypeIsQualityPoor  RejectReasonType = "QUALITY_POOR"
	RejectReasonTypeIsExaggeration RejectReasonType = "EXAGGERATION"
	RejectReasonTypeIsElse         RejectReasonType = "ELSE"
)

func (t RejectReasonType) String() string {
	switch t {
	case RejectReasonTypeIsNone:
		return "无建议"
	case RejectReasonTypeIsReview:
		return "审核不通过"
	case RejectReasonTypeIsLowMaterail:
		return "低质素材"
	case RejectReasonTypeIsDiscomfort:
		return "引人不适"
	case RejectReasonTypeIsQualityPoor:
		return "素材质量低"
	case RejectReasonTypeIsExaggeration:
		return "夸大宣传"
	case RejectReasonTypeIsElse:
		return "其他"
	default:
		return ""
	}
}
