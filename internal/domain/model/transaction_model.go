package model

import (
	"fintrack/internal/infra/zap-logging/log"
	decimal "fintrack/internal/pkg/decimal_clone"
	"fmt"
	"time"
)

type Transaction struct {
	Description string          `json:"description"`
	Timestamp   time.Time       `json:"timestamp"`
	Amount      decimal.Decimal `json:"amount"`
}

type Transactions = []Transaction

func (t Transaction) String() string {
	return fmt.Sprintf("(%v) %-10v %v", t.Timestamp.Format("2006-01-02 15:04:05"), t.Description, t.Amount)
}

func NewTransaction(description string, amount string) Transaction {
	log.Info("New transaction ", description, amount)
	d, err := decimal.NewFromString(amount)
	if err != nil {
		log.Error("Unable to parse amount as decimal", err)
	}
	return Transaction{description, time.Now(), d}
}
