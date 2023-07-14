package tools

import (
	"github.com/antihax/optional"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type ToolAdminGetInfoOpts struct {
	AdvertiserId optional.String
	Codes        optional.Interface
	Language     optional.String
	SubDistrict  optional.String
}

type ToolAdminInfoGetResponse struct {
	common.ResponseBaseInfo
	Data *ToolAdminInfoGetResponseData `json:"data,omitempty"`
}

type ToolAdminInfoGetResponseData struct {
	Districts *[]Districts `json:"districts,omitempty"`
	Version   *string      `json:"version,omitempty"`
}

type Districts struct {
	Name         *string      `json:"name,omitempty"`
	Level        *string      `json:"level,omitempty"`
	Code         *string      `json:"code,omitempty"`
	GeonameId    *uint32      `json:"geoname_id,omitempty"`
	SubDistricts *[]Districts `json:"sub_districts,omitempty"`
}
