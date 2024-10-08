package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequest(res http.ResponseWriter, req *http.Request) {
	slog.Info("request", 
		slog.String("method", req.Method),
		slog.String("url", req.URL.Path),
	)
}
