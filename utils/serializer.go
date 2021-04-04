package utils

import "net/http"

// ResponseResult http 响应结果
type ResponseResult struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// BuildResponse build response
func BuildResponse(code int64, msg string, data interface{}) ResponseResult {
	return ResponseResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// BuildOKResponse Ok response
func BuildOKResponse(data interface{}) ResponseResult {
	// http.StatusOK
	return BuildResponse(0, "success", data)
}

// BuildBadResponse 400 response
func BuildBadResponse(msg string) ResponseResult {
	if msg == "" {
		msg = "请求失败"
	}
	return BuildResponse(http.StatusBadRequest, msg, nil)
}
