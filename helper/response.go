package helper

type Response struct {
	Message interface{} `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildShortResponse(data interface{}) Response {
	res := Response{
		Message: nil,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err interface{}, data interface{}) Response {
	res := Response{
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}

func BuildCustomErrorResponse(message string, err interface{}) Response {
	res := Response{
		Message: message,
		Errors:  err,
		Data:    EmptyObj{},
	}
	return res
}
