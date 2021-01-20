package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping", healthcheck)

	listenAddr := os.Getenv("LISTEN_ADDRESS")
	if listenAddr == "" {
		listenAddr = ":8080"
		log.Printf("No LISTEN_ADDRESS specified. Listening on: %s\n", listenAddr)
	} else {
		log.Printf("Listening on: %s\n", listenAddr)
	}

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
