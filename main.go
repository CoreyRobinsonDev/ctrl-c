package main

import (
	"log/slog"
	"net/http"

	"ctrl-c/routes"
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
		),
	}
	slog.Info("server started on port " + port)
	server.ListenAndServe()
}

