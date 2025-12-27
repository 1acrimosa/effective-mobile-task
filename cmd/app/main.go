package main

import (
	"effective-mobile-task/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	_ "effective-mobile-task/config"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	cfg := config.Load()

	log.Println("server started on :" + cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))

}
