package entity

import (
	decimal "fintrack/internal/pkg/decimal_clone"
	"time"
)

type Transaction struct {
	ID          uint64          `json:"id"`
	UserID      uint64          `json:"user_id"` // Foreign key to User entity
	Year        int             `json:"year"`
	Month       int             `json:"month"`
	Timestamp   time.Time       `json:"timestamp"`
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"` // For handling monetary values with high precision
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
