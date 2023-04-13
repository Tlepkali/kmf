package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"kmf/internal/models"
)

type CurrencyService struct {
	CurrencyRepo *models.CurrencyRepo
}

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Fullname    string  `xml:"fullname"`
	Title       string  `xml:"title"`
	Description float64 `xml:"description"`
}

func NewCurrencyService(repo *models.CurrencyRepo) *CurrencyService {
	return &CurrencyService{
		CurrencyRepo: repo,
	}
}

func (s *CurrencyService) SaveCurrency(date string) error {
	url := fmt.Sprintf("https://nationalbank.kz/rss/get_rates.cfm?fdate=%s", date)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var rates Rates
	err = xml.Unmarshal(body, &rates)
	if err != nil {
		return err
	}

	timeDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return err
	}

	for _, item := range rates.Items {
		if err != nil {
			return err
		}

		c := models.Currency{
			Title: item.Fullname,
			Code:  item.Title,
			Value: item.Description,
			ADate: timeDate,
		}
		go s.CurrencyRepo.Insert(&c)
	}

	return nil
}

func (s *CurrencyService) GetCurrency(date, code string) ([]byte, error) {
	timeDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return nil, err
	}

	if code == "" {
		currencies, err := s.CurrencyRepo.GetByDate(timeDate)
		if err != nil {
			return nil, err
		}
		if len(currencies) == 0 {
			return nil, models.ErrRecordNotFound
		}

		return json.Marshal(currencies)
	}

	currencies, err := s.CurrencyRepo.GetByDateAndCode(code, timeDate)
	if err != nil {
		return nil, err
	}
	if len(currencies) == 0 {
		return nil, models.ErrRecordNotFound
	}

	return json.Marshal(currencies)
}
