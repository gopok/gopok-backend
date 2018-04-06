package core

/*
HttpError is a value that WrapRest can use to set a http code.
*/
type HttpError interface {
	HTTPCode() int
}
