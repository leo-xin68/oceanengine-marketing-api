package tools

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
	model "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/tools"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ToolCountryApiService service

// Info 查询国家/区域信息
func (o *ToolCountryApiService) Info(ctx context.Context, Optionals *model.ToolCountryGetInfoOpts) (model.ToolCountryInfoGetResponseData, http.Header, error) {
	var (
		path        = o.client.Cfg.BasePath + "/2/tools/country/info/"
		queryParams = url.Values{}
		returnValue model.ToolCountryInfoGetResponseData
		response    model.ToolCountryInfoGetResponse
	)

	if Optionals != nil {

		if Optionals.AdvertiserId.IsSet() {
			queryParams.Add("advertiser_id", util.ParameterToString(Optionals.AdvertiserId.Value(), ""))
		}

		if Optionals.Language.IsSet() {
			queryParams.Add("language", util.ParameterToString(Optionals.Language.Value(), ""))
		}
	}

	r, err := o.client.BuildGetDefaultRequest(ctx, path, queryParams)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := o.client.CallApi(r)
	if err != nil || httpResponse == nil {
		return returnValue, nil, err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	defer httpResponse.Body.Close()
	if err != nil {
		return returnValue, nil, err
	}

	if httpResponse.StatusCode < 300 {
		err = o.client.Decode(&response, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *response.Code != 0 {
				var responseErrors []ApiErrorStruct
				if response.Errors != nil {
					responseErrors = *response.Errors
				}
				err = errors.NewError(response.Code, response.Message, responseErrors)
				return returnValue, httpResponse.Header, err
			}
			if response.Data == nil {
				return returnValue, httpResponse.Header, err
			} else {
				return *response.Data, httpResponse.Header, err
			}
		} else {
			return returnValue, httpResponse.Header, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := errors.GenericSwaggerError{}
		newErr.SetBody(responseBody)
		newErr.SetError(httpResponse.Status)
		return returnValue, httpResponse.Header, newErr
	}

	return returnValue, httpResponse.Header, nil
}
