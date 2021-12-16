package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var count int

func handlerResponse(w http.ResponseWriter, r *http.Request) {
	count += 1
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"remote_addres": r.RemoteAddr,
		"requested_url": r.RequestURI,
		"count":         fmt.Sprintf("%d", count),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerResponse)

	http.ListenAndServe(":8000", mux)
}
