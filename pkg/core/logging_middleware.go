package core

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type loggerWriterData struct {
	didSetCode bool
}

type loggerWriter struct {
	innerWriter http.ResponseWriter
	data        *loggerWriterData
}

func (lw loggerWriter) Header() http.Header {
	return lw.innerWriter.Header()
}

func (lw loggerWriter) Write(data []byte) (int, error) {
	log.Print(lw.data)
	if lw.data.didSetCode {
		
	}
	return lw.innerWriter.Write(data)
}
func (lw loggerWriter) WriteHeader(statusCode int) {
	lw.data.didSetCode = true
	lw.innerWriter.WriteHeader(statusCode)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(loggerWriter{w, &loggerWriterData{}}, r)

	})
}
