package enums

// LocationType :受众位置类型
type LocationType string

const (
	HomeLocationType    LocationType = "HOME"
	TravelLocationType  LocationType = "TRAVEL"
	AllLocationType     LocationType = "ALL"
	CurrentLocationType LocationType = "CURRENT"
)

func (t LocationType) String() string {
	switch t {
	case HomeLocationType:
		return "居住在该地区的用户"
	case TravelLocationType:
		return "到该地区旅行的用户"
	case AllLocationType:
		return "该地区内的所有用户"
	case CurrentLocationType:
		return "正在该地区的用户"
	default:
		return ""
	}
}
