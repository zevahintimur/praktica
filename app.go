package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	count      int
	infoLogger *log.Logger
)

func handlerResponse(w http.ResponseWriter, r *http.Request) {
	count += 1
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"remote_addres": r.RemoteAddr,
		"requested_url": r.RequestURI,
		"count":         fmt.Sprintf("%d", count),
	})
	infoLogger.Printf("Get %d request(s). Last from %s", count, r.RemoteAddr)
}

func main() {
	infoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerResponse)
	infoLogger.Println("Starting web server")
	http.ListenAndServe(":8000", mux)
}
