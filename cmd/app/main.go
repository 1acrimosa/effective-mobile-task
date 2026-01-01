package main

import (
	"log"
	"net/http"

	"effective-mobile-task/internal/db"
	"effective-mobile-task/internal/handler"
	"effective-mobile-task/internal/repository"
	"effective-mobile-task/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	dbConn, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewSubscriptionRepository(dbConn)
	service := service.NewSubscriptionService(repo)
	handler := handler.NewSubscriptionHandler(service)

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Get("/subscriptions", handler.GetAll)
	r.Post("/subscriptions", handler.Create)

	log.Println("server started on :8080")
	http.ListenAndServe(":8080", r)
}
