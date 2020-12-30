package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func requestLogger(w http.ResponseWriter, r *http.Request) {
	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(req))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	requestLogger(w, r)
	switch r.Method {
	case http.MethodGet:
		handleVerify(w, r)
	case http.MethodPost:
		handleUpdate(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {

}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	_, _ = w.Write([]byte(challenge))
}

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, http.HandlerFunc(handleRequest))
	if err != nil {
		panic(err)
	}
}
