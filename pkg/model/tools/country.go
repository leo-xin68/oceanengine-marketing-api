package tools

import (
	"github.com/antihax/optional"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type ToolCountryGetInfoOpts struct {
	AdvertiserId optional.String
	Language     optional.String
}

type ToolCountryInfoGetResponse struct {
	common.ResponseBaseInfo
	Data *ToolCountryInfoGetResponseData `json:"data,omitempty"`
}

type ToolCountryInfoGetResponseData struct {
	Districts *[]CountryDistricts `json:"districts,omitempty"`
}

type CountryDistricts struct {
	Code        *string `json:"code,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
