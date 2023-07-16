package enums

// LocationType :受众位置类型
type LocationType string

const (
	LocationTypeIsHome    LocationType = "HOME"
	LocationTypeIsTravel  LocationType = "TRAVEL"
	LocationTypeIsAll     LocationType = "ALL"
	LocationTypeIsCurrent LocationType = "CURRENT"
)

func (t LocationType) String() string {
	switch t {
	case LocationTypeIsHome:
		return "居住在该地区的用户"
	case LocationTypeIsTravel:
		return "到该地区旅行的用户"
	case LocationTypeIsAll:
		return "该地区内的所有用户"
	case LocationTypeIsCurrent:
		return "正在该地区的用户"
	default:
		return ""
	}
}
