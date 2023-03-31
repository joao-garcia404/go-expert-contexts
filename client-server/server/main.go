package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Initialiazing request.")

	defer log.Println("Request done.")

	select {
	case <-time.After(5 * time.Second):
		// Logged at command line stdout
		log.Println("Request processed with success.")

		// Logged at client
		w.Write([]byte("Request processed with success."))
		return

	case <-ctx.Done():
		log.Println("Request cancelled by client")
		return
	}
}
