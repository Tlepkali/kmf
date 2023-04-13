package handlers

import (
	"net/http"

	"kmf/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *service.CurrencyService
}

func NewHandler(s *service.CurrencyService) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/currency/save/{date}", h.saveCurrency).Methods("GET")
	router.HandleFunc("/currency/{date}/{code}", h.getCurrency).Methods("GET")
	router.HandleFunc("/currency/{date}", h.getCurrency).Methods("GET")

	return router
}
