package enums

// BidSpeed : 投放速度
type BidSpeed string

const (
	BidSpeedIsBalance BidSpeed = "BALANCE"
	BidSpeedIsFast    BidSpeed = "FAST"
)

func (t BidSpeed) String() string {
	switch t {
	case BidSpeedIsBalance:
		return "匀速"
	case BidSpeedIsFast:
		return "加速"
	default:
		return ""
	}
}
