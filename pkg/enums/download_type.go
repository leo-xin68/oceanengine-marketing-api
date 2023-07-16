package enums

// DownloadType : 下载方式
type DownloadType string

const (
	DownloadTypeIsDownloadUrl DownloadType = "DOWNLOAD_URL"
	DownloadTypeIsExternalUrl DownloadType = "EXTERNAL_URL"
)

func (t DownloadType) String() string {
	switch t {
	case DownloadTypeIsDownloadUrl:
		return "直接下载"
	case DownloadTypeIsExternalUrl:
		return "落地页下载"
	default:
		return ""
	}
}
