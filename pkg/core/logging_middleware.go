package core

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type loggerWriterData struct {
	didSetCode bool
	statusCode int
	url        string
	method     string
}

type loggerWriter struct {
	innerWriter http.ResponseWriter
	data        *loggerWriterData
}

func (lw loggerWriter) Header() http.Header {
	return lw.innerWriter.Header()
}

func (lw loggerWriter) Write(data []byte) (int, error) {
	if !lw.data.didSetCode {
		lw.data.statusCode = 200
	}
	log.Infof("%s %s %d", lw.data.method, lw.data.url, lw.data.statusCode)
	return lw.innerWriter.Write(data)
}
func (lw loggerWriter) WriteHeader(statusCode int) {
	lw.data.didSetCode = true
	lw.data.statusCode = statusCode
	lw.innerWriter.WriteHeader(statusCode)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(loggerWriter{w, &loggerWriterData{
			method: r.Method,
			url:    r.RequestURI,
		}}, r)

	})
}
