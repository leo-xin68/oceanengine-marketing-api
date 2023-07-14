package enums

// LandingPageStayTime : 店铺停留时长,单位为毫秒
type LandingPageStayTime uint16

const (
	LandingPageStayTime1000  LandingPageStayTime = 1000
	LandingPageStayTime5000  LandingPageStayTime = 5000
	LandingPageStayTime12000 LandingPageStayTime = 12000
	LandingPageStayTime20000 LandingPageStayTime = 20000
	LandingPageStayTime30000 LandingPageStayTime = 30000
	LandingPageStayTime40000 LandingPageStayTime = 40000
	LandingPageStayTime50000 LandingPageStayTime = 50000
	LandingPageStayTime60000 LandingPageStayTime = 60000
)

func (t LandingPageStayTime) String() string {
	switch t {
	case LandingPageStayTime1000:
		return "1000"
	case LandingPageStayTime5000:
		return "5000"
	case LandingPageStayTime12000:
		return "12000"
	case LandingPageStayTime20000:
		return "20000"
	case LandingPageStayTime30000:
		return "30000"
	case LandingPageStayTime40000:
		return "40000"
	case LandingPageStayTime50000:
		return "50000"
	case LandingPageStayTime60000:
		return "60000"
	default:
		return "0"
	}
}
