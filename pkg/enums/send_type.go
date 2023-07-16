package enums

// SendType : 数据发送方式
type SendType string

const (
	SendTypeIsServer SendType = "SERVER_SEND"
	SendTypeIsClient SendType = "CLIENT_SEND"
)

func (t SendType) String() string {
	switch t {
	case SendTypeIsServer:
		return "服务器端上传"
	case SendTypeIsClient:
		return "客户端上传"
	default:
		return ""
	}
}
