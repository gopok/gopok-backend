package core

/*
ErrorResponse stores error responses in json format
*/

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string        `json:"message"`
}
func (e ErrorResponse) HTTPCode() int {
	return e.Code
}


type HttpError interface {
	HTTPCode() int
}

func NewErrorResponse(message string, code int) ErrorResponse {
	return ErrorResponse{code, message}
}