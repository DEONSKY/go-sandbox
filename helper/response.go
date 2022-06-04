package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	spilledError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  spilledError,
		Data:    data,
	}
	return res
}

func BuildCustomErrorResponse(message string, err string) Response {
	res := Response{
		Status:  false,
		Message: message,
		Errors:  err,
		Data:    EmptyObj{},
	}
	return res
}
