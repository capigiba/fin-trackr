package currency

import (
	"database/sql"
	"time"
)

type CurrencyRepository interface {
	SaveConversionRate(from, to string, rate float32, date time.Time) error
	GetConversionRate(from, to string, date time.Time) (float32, error)
}

type currencyRepositoryImpl struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) CurrencyRepository {
	return &currencyRepositoryImpl{db: db}
}

func (r *currencyRepositoryImpl) SaveConversionRate(from, to string, rate float32, date time.Time) error {
	_, err := r.db.Exec("INSERT INTO conversions (from_currency, to_currency, rate, date) VALUES (?, ?, ?, ?)", from, to, rate, date)
	return err
}

func (r *currencyRepositoryImpl) GetConversionRate(from, to string, date time.Time) (float32, error) {
	var rate float32
	err := r.db.QueryRow("SELECT rate FROM conversions WHERE from_currency = ? AND to_currency = ? AND date = ?", from, to, date).Scan(&rate)
	return rate, err
}
