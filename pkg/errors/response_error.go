package errors

import . "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/common"

// ResponseError ...
type ResponseError struct {
	error
	Code      int64            `json:"code,omitempty"`
	Message   string           `json:"message,omitempty"`
	MessageCn string           `json:"message_cn,omitempty"`
	Errors    []ApiErrorStruct `json:"errors,omitempty"`
}

// Error ...
func (e ResponseError) Error() string {
	return e.Message
}

// NewError ...
func NewError(code *int64, message *string, errors []ApiErrorStruct) ResponseError {
	var codeValue int64
	var messageValue string
	if code != nil {
		codeValue = *code
	}
	if message != nil {
		messageValue = *message
	}

	return ResponseError{
		Code:    codeValue,
		Message: messageValue,
		Errors:  errors,
	}
}
