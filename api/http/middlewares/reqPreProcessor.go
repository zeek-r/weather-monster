package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zeek-r/weather-monster/api/http/utils/iptool"
	"github.com/zeek-r/weather-monster/pkg/logger"
)

func TagRequestID(r *http.Request) {
	r.Header.Add("request_id", uuid.New().String())
}

func logRequest(r *http.Request) {
	logger.Info("Incoming Request", logger.Fields{
		"path":       r.URL.Path,
		"ip_address": iptool.GetIPAddressFromRequest(r),
		"method":     r.Method,
		"request_id": r.Header.Get("request_id"),
	})
}

func RequestPreProcessor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TagRequestID(r)
		logRequest(r)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
