package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	apierrors "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

// ApiClient 负责请求api。在大多数情况下，应该只有一个共享的APIClient
type ApiClient struct {
	Cfg       *config.Configuration
	SdkConfig *config.SdkConfig
	common    service // 复用单个结构，而不是为每个服务分配一个结构

	// API Services

	ProjectApi *ProjectApiService
}

type service struct {
	client *ApiClient
}

// NewApiClient 创建新的API客户端。需要描述应用程序的userAgent字符串。
// 可选的自定义http.Client，来支持缓存之类的高级功能。
func NewApiClient(sdkConfig *config.SdkConfig) *ApiClient {
	cfg := sdkConfig.Configuration
	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{}
	}

	c := &ApiClient{}
	c.SdkConfig = sdkConfig
	c.Cfg = &cfg
	c.common.client = c

	// API Services
	c.ProjectApi = (*ProjectApiService)(&c.common)

	return c
}

// callApi 发起请求.
func (c *ApiClient) callApi(request *http.Request) (*http.Response, error) {
	return c.Cfg.HttpClient.Do(request)
}

// buildRequest 构建GET请求
func (c *ApiClient) buildGetRequest(ctx context.Context, path string, body interface{}) (req *http.Request, err error) {
	httpMethod := strings.ToUpper("Get")
	var (
		queryParams  = url.Values{}
		formParams   = url.Values{}
		headerParams = make(map[string]string)
		fileName     string
		fileBytes    []byte
		fileKey      string
	)
	headerParams["Content-Type"] = "application/json"
	req, err = c.buildRequest(ctx, path, httpMethod, body, headerParams, queryParams, formParams, fileName, fileBytes, fileKey)
	return req, err
}

// buildRequest 构建POST请求
func (c *ApiClient) buildPostRequest(ctx context.Context, path string, body interface{}) (req *http.Request, err error) {
	httpMethod := strings.ToUpper("Post")
	var (
		queryParams  = url.Values{}
		formParams   = url.Values{}
		headerParams = make(map[string]string)
		fileName     string
		fileBytes    []byte
		fileKey      string
	)
	headerParams["Content-Type"] = "application/json"
	req, err = c.buildRequest(ctx, path, httpMethod, body, headerParams, queryParams, formParams, fileName, fileBytes, fileKey)
	return req, err
}

// buildRequest 构建请求
func (c *ApiClient) buildRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string, fileBytes []byte, fileKey string,
) (req *http.Request, err error) {

	var body *bytes.Buffer

	// post请求body
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// 表单请求
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					_ = w.WriteField(k, iv)
				}
			}
		}

		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(fileKey, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		_ = w.Close()

	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	//if strings.HasPrefix(headerParams["Content-Type"], "application/json") && len(queryParams) > 0 {
	//	var tmp bytes.Buffer
	//	jsonParams, _ := json.Marshal(queryParams)
	//	tmp.WriteString("?")
	//	tmp.WriteString(string(jsonParams))
	//	queryParams = make(map[string][]string)
	//}

	// 设置路径和查询参数
	reqUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := reqUrl.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode参数
	reqUrl.RawQuery = query.Encode()

	// 生成请求
	if body != nil {
		req, err = http.NewRequest(method, reqUrl.String(), body)
	} else {
		req, err = http.NewRequest(method, reqUrl.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// 添加header参数
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		req.Header = headers
	}

	for header, value := range c.Cfg.DefaultHeader {
		req.Header.Add(header, value)
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func (c *ApiClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/json") {
		dec := json.NewDecoder(bytes.NewReader(b))
		if c.SdkConfig.IsStrictMode {
			dec.DisallowUnknownFields() // Force
		}
		if err := dec.Decode(v); err != nil {
			return apierrors.NewResponseStrictError(err.Error())
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// parameterToString 将obi 参数转换为字符串，如果提供了格式，则使用分隔符。
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	case "multi":
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}

	return fmt.Sprintf("%v", obj)
}

// GenericSwaggerError 报错的信息结构
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

func (e GenericSwaggerError) Error() string {
	return e.error
}

func (e GenericSwaggerError) Body() []byte {
	return e.body
}

func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
