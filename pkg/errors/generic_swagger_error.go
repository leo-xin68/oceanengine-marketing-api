package errors

// GenericSwaggerError 报错的信息结构
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

func (e GenericSwaggerError) SetBody(body []byte) {
	e.body = body
}
func (e GenericSwaggerError) SetError(err string) {
	e.error = err
}
func (e GenericSwaggerError) SetModel(model any) {
	e.model = model
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
