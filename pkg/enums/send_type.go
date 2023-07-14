package enums

// SendType : 数据发送方式
type SendType string

const (
	ServerSendType SendType = "SERVER_SEND"
	ClientSendType SendType = "CLIENT_SEND"
)

func (t SendType) String() string {
	switch t {
	case ServerSendType:
		return "服务器端上传"
	case ClientSendType:
		return "客户端上传"
	default:
		return ""
	}
}
