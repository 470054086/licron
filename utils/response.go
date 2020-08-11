package utils

import (
	"net/http"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

//  验证格式参数格式错误的提示
func ResponseParamError() *Response {
	r := &Response{Message: "参数格式错误", Code: http.StatusBadRequest, Data: make(map[string]interface{})}
	return r
}

func ResponseMessageError(msg string) *Response {
	r := &Response{Message: msg, Code: http.StatusBadRequest, Data: make(map[string]interface{})}
	return r
}


func ResponseActionError() *Response {
	r := &Response{Message: "操作失败", Code: http.StatusBadRequest, Data: make(map[string]interface{})}
	return r
}

func ResponseSuccess() *Response {
	r := &Response{Message: "操作成功", Code: 0, Data: make(map[string]interface{})}
	return r
}
