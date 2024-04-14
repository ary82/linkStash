package logging

import "net/http"

type customWriter struct {
	http.ResponseWriter
	statuscode int
}

func (w *customWriter) WriteHeader(statuscode int) {
	w.ResponseWriter.WriteHeader(statuscode)
	w.statuscode = statuscode
}
