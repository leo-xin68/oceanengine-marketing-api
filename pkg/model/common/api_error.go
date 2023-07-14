package common

type ApiErrorStruct struct {
	Type_                    *string                   `json:"type,omitempty"`
	CommonParamExceptionSpec *CommonParamExceptionSpec `json:"common_param_exception_spec,omitempty"`
}

type CommonParamExceptionSpec struct {
	ErrorFields *[]string `json:"error_fields,omitempty"`
	ErrorMsg    *string   `json:"error_msg,omitempty"`
}
