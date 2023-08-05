package common

import "encoding/json"

type DSWError struct {
	Code    int
	Message any
}

func NewDSWError(code int, message any) *DSWError {
	return &DSWError{Code: code, Message: message}
}

func (D *DSWError) Error() string {
	err := NewDSWError(D.Code, D.Message)
	res, _ := json.Marshal(err)
	return string(res)
}

func (D *DSWError) RuntimeError() {
	//TODO implement me
	panic("implement me")
}
