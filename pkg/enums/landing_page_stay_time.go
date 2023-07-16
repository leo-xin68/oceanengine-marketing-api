package enums

// LandingPageStayTime : 店铺停留时长,单位为毫秒
type LandingPageStayTime uint16

const (
	LandingPageStayTimeIs1000  LandingPageStayTime = 1000
	LandingPageStayTimeIs5000  LandingPageStayTime = 5000
	LandingPageStayTimeIs12000 LandingPageStayTime = 12000
	LandingPageStayTimeIs20000 LandingPageStayTime = 20000
	LandingPageStayTimeIs30000 LandingPageStayTime = 30000
	LandingPageStayTimeIs40000 LandingPageStayTime = 40000
	LandingPageStayTimeIs50000 LandingPageStayTime = 50000
	LandingPageStayTimeIs60000 LandingPageStayTime = 60000
)

func (t LandingPageStayTime) String() string {
	switch t {
	case LandingPageStayTimeIs1000:
		return "1000"
	case LandingPageStayTimeIs5000:
		return "5000"
	case LandingPageStayTimeIs12000:
		return "12000"
	case LandingPageStayTimeIs20000:
		return "20000"
	case LandingPageStayTimeIs30000:
		return "30000"
	case LandingPageStayTimeIs40000:
		return "40000"
	case LandingPageStayTimeIs50000:
		return "50000"
	case LandingPageStayTimeIs60000:
		return "60000"
	default:
		return "0"
	}
}
