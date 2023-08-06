package common

type ResCommonHandler struct {
	Code    int
	Message any
}

func (r *ResCommonHandler) HandRes() any {

	return NewDSWHandler(r.Code, r.Message)
}

func NewDSWHandler(code int, message any) *ResCommonHandler {
	return &ResCommonHandler{Code: code, Message: message}
}

type Handler interface {
	HandRes() any
}
