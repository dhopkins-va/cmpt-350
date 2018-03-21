package main

import (
	"net/http"
	"time"
	"encoding/json"
	"fmt"
)

type TimeResponse struct {
	Time time.Time `json:"time"`
	Server string `json:"server"`
}

func main() {
	fmt.Printf("Server Started\n")
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TimeResponse{Time: time.Now(), Server:"Container"})
		fmt.Printf("* Request Handled")
	})
	http.ListenAndServe(":80", nil)
	fmt.Printf("Server Terminated\n")
}
