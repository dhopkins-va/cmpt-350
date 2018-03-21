package main

import (
	"net/http"
	"time"
	"encoding/json"
)

type TimeResponse struct {
	Time time.Time `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TimeResponse{Time: time.Now()})
	})
	http.ListenAndServe(":80", nil)
}
