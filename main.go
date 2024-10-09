package main

import (
	"ctrl-c/middleware"
	"ctrl-c/routes"
	"log/slog"
	"net/http"

	u "ctrl-c/util"

	_ "github.com/lib/pq"
)


func main() {
	router := http.NewServeMux()
	port := u.Unwrap(u.Dotenv("PORT"))

	router.HandleFunc("GET /health", routes.Health)

	server := http.Server {
		Addr: port, 
		Handler: u.SetMiddleware(router,
			middleware.LogRequest,
		),
	}
	slog.Info("server started on port " + port)
	server.ListenAndServe()
}

