package util

import (
	"bufio"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func Dotenv(variable string) (string, error) {
	file := Unwrap(os.Open(".env"))
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		entry := strings.Split(reader.Text(), "=")
		if len(entry) != 2 {
			return "", errors.New("malformed .env file")
		}
		if entry[0] == variable {
			return entry[1], nil
		}
	}

	return "", errors.New("variable " + variable + " does not exist")
}


func SetMiddleware(router http.Handler, middlewares ...http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer router.ServeHTTP(res, req)
		for _, mw := range middlewares {
			mw(res, req)
		}
	}
}

func Expect(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func Unwrap[T any](val T, err error) T {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return val
}

func UnwrapOr[T any](val T, err error) func(T) T {
	if err != nil {
		slog.Error(err.Error())
		return func(d T) T {
			return d
		}
	} else {
		return func(_ T) T {
			return val
		}
	}
}

func UnwrapOrElse[T any](val T, err error) func(func() T) T {
	if err != nil {
		slog.Error(err.Error())
		return func(fn func() T) T {
			return fn()
		}
	} else {
		return func(_ func() T) T {
			return val
		}
	}

}
