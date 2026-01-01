package handler

import (
	"effective-mobile-task/internal/service"
	"encoding/json"
	"net/http"
)

type SubscriptionService interface {
	GetAll() (interface{}, error)
}

type SubscriptionHandler struct {
	service SubscriptionService
}

func NewSubscriptionHandler(service *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

func (h *SubscriptionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	subs, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

func (h *SubscriptionHandler) Create(writer http.ResponseWriter, request *http.Request) {

}
