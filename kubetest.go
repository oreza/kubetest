package main

import (
"log"
"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("Hello form 1 & 0 \n"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("I am healthy \n"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}