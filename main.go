package main

import (
	"ctrl-c/database"
	"ctrl-c/middleware"
	"log/slog"
	"net/http"

	u "ctrl-c/util"

	_ "github.com/lib/pq"
)


func main() {
	db := database.Open()
	defer db.Close()
	router := http.NewServeMux()
	port := u.Unwrap(u.Dotenv("PORT"))

	router.HandleFunc("GET /users/{userId}", test)

	server := http.Server {
		Addr: port, 
		Handler: u.SetMiddleware(router,
			middleware.LogRequest,
		),
	}
	slog.Info("server started on port " + port)
	server.ListenAndServe()
}



func test(res http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")

	res.Write([]byte("User Id: " + userId))
}
