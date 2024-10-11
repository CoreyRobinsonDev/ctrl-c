package routes

import (
	"ctrl-c/database"
	"ctrl-c/util"
	"encoding/json"
	"net/http"
	"time"
)

type health struct {
	Status string `json:"status"`
	Downstream string `json:"downstream"`
	Timestamp time.Time `json:"timestamp"`
}

type db struct {
	Status
}
func Health(res http.ResponseWriter, req *http.Request) {
	status := health{"UP", "", time.Now()}
	db := database.Open()
	defer db.Close()

	if err := db.Ping(); err != nil {
		status.Downstream = "DOWN"
	} else {
		status.Downstream = "UP"
	}
	bytes := util.Unwrap(json.Marshal(status))
	res.Header().Set("Content-Type", "application/json")
	res.Write(bytes)
}
