package api

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"
	"io/ioutil"
	"net/http"
)

type PromotionApiService service

// List 获取广告列表
func (o *PromotionApiService) List(ctx context.Context, data model.PromotionGetRequest) (model.PromotionGetResponseData, http.Header, error) {
	var (
		path        = o.client.Cfg.BasePath + "/v3.0/promotion/list/"
		returnValue model.PromotionGetResponseData
		response    model.PromotionGetResponse
	)

	r, err := o.client.BuildGetRequest(ctx, path, data)
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
		newErr.SetModel(data)
		newErr.SetBody(responseBody)
		newErr.SetError(httpResponse.Status)
		return returnValue, httpResponse.Header, newErr
	}

	return returnValue, httpResponse.Header, nil
}

// Create 创建广告
func (o *PromotionApiService) Create(ctx context.Context, data model.PromotionCreateRequest) (model.PromotionCreateResponseData, http.Header, error) {
	var (
		path        = o.client.Cfg.BasePath + "/v3.0/promotion/create/"
		returnValue model.PromotionCreateResponseData
		response    model.PromotionCreateResponse
	)

	r, err := o.client.BuildPostRequest(ctx, path, data)
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
