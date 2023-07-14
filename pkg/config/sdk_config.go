package config

import "net/http"

// SdkConfig ...
type SdkConfig struct {
	Configuration
	AppID        uint64
	AccessToken  string
	IsDebug      bool
	IsStrictMode bool
	GlobalConfig GlobalConfig
}

func NewSdkConfig(AppId uint64, AccessToken string, IsDebug bool) *SdkConfig {
	cfg := &SdkConfig{
		AppID:       AppId,
		AccessToken: AccessToken,
		IsDebug:     IsDebug,
		Configuration: Configuration{
			BasePath: "https://ad.oceanengine.com/open_api",
			Version:  "1.0.0",
		},
	}
	return cfg
}

// GlobalConfig ...
type GlobalConfig struct {
	HttpOption HttpOption
}

// HttpOption ...
type HttpOption struct {
	Header http.Header
}
