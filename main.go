package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeStamp struct {
	Time string `json:"time"`
}

const port = ":3001"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	now := time.Now().Format("20060102150405")
	timeStamp := TimeStamp{Time: now}
	jsonString, err := json.Marshal(timeStamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonString)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on %s...", port)
	http.ListenAndServe(port, nil)
}
