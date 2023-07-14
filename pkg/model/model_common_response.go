package model

type ResponsePageInfo struct {
	Page        *int64 `json:"page,omitempty"`
	PageSize    *int64 `json:"page_size,omitempty"`
	TotalNumber *int64 `json:"total_number,omitempty"`
	TotalPage   *int64 `json:"total_page,omitempty"`
}

type ResponseBaseInfo struct {
	Code      *int64            `json:"code,omitempty"`
	Message   *string           `json:"message,omitempty"`
	RequestId *string           `json:"request_id,omitempty"`
	Errors    *[]ApiErrorStruct `json:"errors,omitempty"`
}
