package main

import (
	"log"
	"net/http"
)

func main() {

	// TODO: INIT CONFIG

	// TODO: INIT LOGGER

	// TODO: INIT STORAGE

	// TODO: init router

	// TODO: run server

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("server started on :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
