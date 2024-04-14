package logging

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Custom writer for getting statuscodes
		writer := &customWriter{
			ResponseWriter: w,
			statuscode:     http.StatusOK,
		}

		start := time.Now()
		next.ServeHTTP(writer, r)
		log.Println(writer.statuscode, r.Method, r.URL.Path, time.Since(start))
	})
}
