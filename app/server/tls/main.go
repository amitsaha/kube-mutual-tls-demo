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
		listenAddr = ":8443"
		log.Printf("No LISTEN_ADDRESS specified. Listening on: %s\n", listenAddr)
	} else {
		log.Printf("Listening on: %s\n", listenAddr)
	}
	TLS_CERT_PATH := os.Getenv("TLS_CERT_PATH")
	TLS_KEY_PATH := os.Getenv("TLS_KEY_PATH")

	if len(TLS_CERT_PATH) == 0 || len(TLS_KEY_PATH) == 0 {
		log.Fatal("Must specify both TLS_CERT_PATH and TLS_KEY_PATH")
	}
	log.Fatal(http.ListenAndServeTLS(listenAddr, TLS_CERT_PATH, TLS_KEY_PATH, nil))
}
