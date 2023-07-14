package model

type CommonParamExceptionSpec struct {
	ErrorFields *[]string `json:"error_fields,omitempty"`
	ErrorMsg    *string   `json:"error_msg,omitempty"`
}
