package main

import (
	"ctrl-c/middleware"
	"flag"
	"log/slog"
	"net/http"
	"os"
)


func main() {
	router := http.NewServeMux()
	port := flag.String("port", ":1337", "server listening port")
	flag.Parse()

	router.HandleFunc("GET /users/{userId}", test)

	server := http.Server {
		Addr: *port, 
		Handler: SetMiddleware(router,
			middleware.LogRequest,
		),
	}
	slog.Info("server started on port " + *port)
	server.ListenAndServe()
}

func SetMiddleware(router http.Handler, middlewares ...http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer router.ServeHTTP(res, req)
		for _, mw := range middlewares {
			mw(res, req)
		}
	}
}

func Expect[T any](val T, err error) T {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return val
}

func test(res http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")

	res.Write([]byte("User Id: " + userId))
}
