package enums

// DownloadType : 下载方式
type DownloadType string

const (
	DownloadUrl DownloadType = "DOWNLOAD_URL"
	ExternalUrl DownloadType = "EXTERNAL_URL"
)

func (t DownloadType) String() string {
	switch t {
	case DownloadUrl:
		return "直接下载"
	case ExternalUrl:
		return "落地页下载"
	default:
		return ""
	}
}
