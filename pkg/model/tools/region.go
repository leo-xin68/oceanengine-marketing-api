package tools

import (
	"github.com/antihax/optional"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
)

type ToolRegionGetInfoOpts struct {
	RegionType  optional.String
	RegionLevel optional.String
}

type ToolRegionInfoGetResponse struct {
	common.ResponseBaseInfo
	Data *ToolRegionInfoGetResponseData `json:"data,omitempty"`
}

type ToolRegionInfoGetResponseData struct {
	List *[]RegionList `json:"list,omitempty"`
}

type RegionList struct {
	Id          *uint32 `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ParentId    *uint32 `json:"parent_id,omitempty"`
	RegionLevel *string `json:"region_level,omitempty"`
}
