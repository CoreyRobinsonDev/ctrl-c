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
	Downstream struct {
		Status string `json:"status"`
	} `json:"downstream"`
	Timestamp time.Time `json:"timestamp"`
}


func Health(res http.ResponseWriter, req *http.Request) {
	resObj := new(health)
	db := database.Open()
	defer db.Close()

	resObj.Status = "UP"
	resObj.Timestamp = time.Now()

	if err := db.Ping(); err != nil {
		resObj.Downstream.Status = "DOWN"
	} else {
		resObj.Downstream.Status = "UP"
	}

	bytes := util.Unwrap(json.Marshal(resObj))
	res.Header().Set("Content-Type", "application/json")
	res.Write(bytes)
}
