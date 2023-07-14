package enums

// BidSpeed : 投放速度
type BidSpeed string

const (
	BalanceBidSpeed BidSpeed = "BALANCE"
	FastBidSpeed    BidSpeed = "FAST"
)

func (t BidSpeed) String() string {
	switch t {
	case BalanceBidSpeed:
		return "匀速"
	case FastBidSpeed:
		return "加速"
	default:
		return ""
	}
}
