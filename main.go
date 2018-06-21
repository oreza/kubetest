package main

import (
"log"
"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("hello form 1 & 0 \n"))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}