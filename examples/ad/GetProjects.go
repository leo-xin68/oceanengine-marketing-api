package main

import (
	"encoding/json"
	"fmt"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/examples"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/ads"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/util"
	"net/http"
	"strings"
)

type ProjectsGetExample struct {
	OceanAds          *ads.SdkClient
	ProjectGetRequest model.ProjectGetRequest
}

func (e *ProjectsGetExample) Init() {
	e.OceanAds = ads.Init(&config.SdkConfig{
		AppID:       AppId,
		AccessToken: AccessToken,
		IsDebug:     true,
	})
}

func (e *ProjectsGetExample) RunExample() (model.ProjectGetResponseData, http.Header, error) {
	oceanAds := e.OceanAds
	// change ctx as needed
	ctx := *oceanAds.Ctx
	return oceanAds.Project().Get(ctx, e.ProjectGetRequest)
}

func main() {
	e := &ProjectsGetExample{}
	e.Init()

	requestJsonDoc := strings.Replace(`
		{
			"advertiser_id":{advertiser_id},
			"page":1,
			"page_size":100
		}
	`, "{advertiser_id}", AdvertiserId, -1)

	json.Unmarshal([]byte(requestJsonDoc), &e.ProjectGetRequest)

	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}

	util.JsonFormatPrint("Request data", e.ProjectGetRequest)
	util.JsonFormatPrint("Response data", response)
	util.JsonPrint("Headers", headers)
}
