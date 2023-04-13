package handlers

import (
	"encoding/json"
	"net/http"

	_ "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"

	"github.com/gorilla/mux"
)

// swagger: response JSONError
type JSONError struct {
	Succes string `json:"success"`
	Error  string `json:"error,omitempty"`
}

// @ Summary Save currency in database
// @ Description Save currency in database from national bank of Kazakhstan API
// @ Tags currency
// @ Accept json from API by date
// @ Param date path string true "Date in format dd.mm.yyyy"
// @ Success 200 {object} JSONError
// @ Failure 500 {object} JSONError
// @ Router /currency/save/{date} [get]
func (h *Handler) saveCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	err := h.Service.SaveCurrency(date)
	var jsonError JSONError
	if err != nil {
		jsonError.Succes = "false"
		jsonError.Error = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonError.Succes = "true"
	json.NewEncoder(w).Encode(jsonError)
}

// @ Summary Get currency from database
// @ Description Get currency from database by date and code, if code is empty only by date
// @ Tags currency
// @ Accept date and code
// @ Param date path string true "Date in format dd.mm.yyyy"
// @ Param code path string false "Code of currency"
// @ Success 200 {object} models.Currency
// @ Failure 500 {object} JSONError
// @ Router /currency/{date}/{code} [get]
func (h *Handler) getCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	code := vars["code"]

	currencies, err := h.Service.GetCurrency(date, code)
	var jsonError JSONError
	if err != nil {
		jsonError.Succes = "false"
		jsonError.Error = "Currency not found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencies)
}
