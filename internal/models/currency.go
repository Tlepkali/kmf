package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Currency struct {
	Title string    `json:"title"`
	Code  string    `json:"code"`
	Value float64   `json:"value"`
	ADate time.Time `json:"a_date"`
}

type CurrencyRepo struct {
	db *sql.DB
}

var ErrRecordNotFound = errors.New("no data for this date or code")

func NewCurrencyRepo(db *sql.DB) *CurrencyRepo {
	return &CurrencyRepo{
		db: db,
	}
}

func (r *CurrencyRepo) Insert(c *Currency) {
	_, err := r.db.Exec("INSERT INTO R_CURRENCY (title, code, value, a_date) VALUES (?, ?, ?, ?)", c.Title, c.Code, c.Value, c.ADate)
	if err != nil {
		log.Println(err)
	}
}

func (r *CurrencyRepo) GetByDate(date time.Time) ([]*Currency, error) {
	rows, err := r.db.Query("SELECT title, code, value, a_date FROM R_CURRENCY  WHERE a_date = ?", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []*Currency
	for rows.Next() {
		var c Currency
		err := rows.Scan(&c.Title, &c.Code, &c.Value, &c.ADate)
		if err != nil {
			return nil, err
		}
		fmt.Println(c)
		currencies = append(currencies, &c)
	}

	return currencies, nil
}

func (r *CurrencyRepo) GetByDateAndCode(code string, date time.Time) ([]*Currency, error) {
	rows, err := r.db.Query("SELECT title, code, value, a_date FROM R_CURRENCY  WHERE code = ? AND a_date = ?", code, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []*Currency
	for rows.Next() {
		var c Currency
		err := rows.Scan(&c.Title, &c.Code, &c.Value, &c.ADate)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, &c)
	}

	return currencies, nil
}
