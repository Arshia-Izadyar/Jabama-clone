package helper

import "github.com/Arshia-Izadyar/Jabama-clone/src/api/validators"

type Response struct {
	Result     interface{}
	StatusCode int
	Success    bool
	Err        string
}

func GenerateResponse(result interface{}, code int, success bool) *Response {
	return &Response{
		Result:     result,
		StatusCode: code,
		Success:    success,
		Err:        "",
	}
}

func GenerateResponseWithError(result interface{}, code int, success bool, err error) *Response {
	return &Response{
		Result:     result,
		StatusCode: code,
		Success:    success,
		Err:        err.Error(),
	}
}

func GenerateResponseWithErrorWithValidationError(code int, success bool, err error) *Response {
	ve := validators.GetValidationError(err)
	return &Response{
		Result:     ve,
		StatusCode: code,
		Success:    success,
		Err:        err.Error(),
	}
}
