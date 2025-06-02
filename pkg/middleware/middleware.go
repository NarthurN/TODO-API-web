package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NarthurN/TODO-API-web/pkg/loger"
)

var HelperNow = time.Now
var HelperSince = time.Since

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Создаём обёртку поверх оригинального ResponseWriter
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // значение по умолчанию
		}

		start := HelperNow()
		next.ServeHTTP(w, r)
		duration := HelperSince(start)

		loger.L.Info(fmt.Sprintf(
			"[%s] %s %s %d %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			lrw.statusCode,
			duration,
		))
	})
}
