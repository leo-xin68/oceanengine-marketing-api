package model

type ApiErrorStruct struct {
	Type_                    *string                   `json:"type,omitempty"`
	CommonParamExceptionSpec *CommonParamExceptionSpec `json:"common_param_exception_spec,omitempty"`
}
