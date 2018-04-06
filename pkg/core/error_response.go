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




func NewErrorResponse(message string, code int) ErrorResponse {
	return ErrorResponse{code, message}
}
