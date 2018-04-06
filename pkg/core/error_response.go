package core

/*
ErrorResponse stores error responses in json format
*/
type ErrorResponse struct {
	Code    int16         `json:"code"`
	Message string        `json:"message"`
}