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

// GlobalConfig ...
type GlobalConfig struct {
	HttpOption HttpOption
}

// HttpOption ...
type HttpOption struct {
	Header http.Header
}
