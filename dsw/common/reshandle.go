package common

type ResCommonHandler struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}
type ResWithData struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	ResCommonHandler
}

func (r *ResCommonHandler) HandRes() any {

	return NewDSWHandler(r.Code, r.Message)
}

func NewDSWHandler(code int, message any) Handler {
	return &ResCommonHandler{Code: code, Message: message}
}
func NewResWithData(Page int, Size int, Code int, Data any, Total int) Handler {
	data := &ResWithData{Size: Size, Page: Page, Total: Total}
	data.Code = Code
	data.Message = Data
	return data
}

type Handler interface {
	HandRes() any
}
