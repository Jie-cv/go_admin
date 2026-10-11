package utils

import "github.com/Jie-cv/go_admin/contants/code"

type Response struct {
	Code code.RespCode `json:"code"`
	Msg  string        `json:"message"`
}

type ErrorResponce struct {
	Response
}

type SuccessResponse[T any] struct {
	Response
	Data T `json:"data"`
}

func FormateSuccess[T any](code code.RespCode, msg string, data T) SuccessResponse[T] {
	return SuccessResponse[T]{Data: data, Code: code, Msg: msg}
}
func FormateError(code code.RespCode, msg string) ErrorResponce {
	return ErrorResponce{Code: code, Msg: msg}
}
