package app

import (
	"database/sql"
	"log"
	"net/http"

	"kmf/internal/handlers"
	"kmf/internal/models"
	"kmf/internal/service"
)

func Run(addr string, db *sql.DB) {
	repo := models.NewCurrencyRepo(db)

	service := service.NewCurrencyService(repo)

	handler := handlers.NewHandler(service)

	if err := http.ListenAndServe(addr, handler.Routes()); err != nil {
		log.Fatal(err)
	}
}
