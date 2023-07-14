package api

import (
	"context"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	. "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model"
	"io/ioutil"
	"net/http"
)

type ProjectApiService service

// Get 获取项目列表
func (o *ProjectApiService) Get(ctx context.Context, data ProjectGetRequest) (ProjectGetResponseData, http.Header, error) {
	var (
		path        = o.client.Cfg.BasePath + "/v3.0/project/list/"
		returnValue ProjectGetResponseData
		response    ProjectGetResponse
	)

	r, err := o.client.buildGetRequest(ctx, path, data)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := o.client.callApi(r)
	if err != nil || httpResponse == nil {
		return returnValue, nil, err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	defer httpResponse.Body.Close()
	if err != nil {
		return returnValue, nil, err
	}

	if httpResponse.StatusCode < 300 {
		err = o.client.decode(&response, responseBody, httpResponse.Header.Get("Content-Type"))
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
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return returnValue, httpResponse.Header, newErr
	}

	return returnValue, httpResponse.Header, nil
}

// Create 创建项目列表
func (o *ProjectApiService) Create(ctx context.Context, data ProjectCreateRequest) (ProjectCreateResponseData, http.Header, error) {
	var (
		path        = o.client.Cfg.BasePath + "/v3.0/project/create/"
		returnValue ProjectCreateResponseData
		response    ProjectCreateResponse
	)

	r, err := o.client.buildPostRequest(ctx, path, data)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := o.client.callApi(r)
	if err != nil || httpResponse == nil {
		return returnValue, nil, err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	defer httpResponse.Body.Close()
	if err != nil {
		return returnValue, nil, err
	}

	if httpResponse.StatusCode < 300 {
		err = o.client.decode(&response, responseBody, httpResponse.Header.Get("Content-Type"))
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
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return returnValue, httpResponse.Header, newErr
	}

	return returnValue, httpResponse.Header, nil
}
