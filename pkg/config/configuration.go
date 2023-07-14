package config

import (
	"net/http"
)

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	HttpClient    *http.Client
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
