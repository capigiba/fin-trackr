package transaction

import (
	"database/sql"
	"fintrack/internal/domain/entity"
	"fintrack/internal/domain/model"
	"fintrack/internal/infra/zap-logging/log"
	decimal "fintrack/internal/pkg/decimal_clone"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) Save(t entity.Transaction) error {
	query := `
	INSERT INTO transactions (user_id, year, month, timestamp, description, amount, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := repo.db.Exec(query, t.UserID, t.Year, t.Month, t.Timestamp, t.Description, t.Amount.StringFixed(2), t.CreatedAt, t.UpdatedAt)
	return err
}

func (repo *TransactionRepository) Load(userID uint64, year, month int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	query := `
    SELECT timestamp, description, amount
    FROM transactions
    WHERE user_id = ? AND year = ? AND month = ?
    ORDER BY timestamp ASC
    `
	rows, err := repo.db.Query(query, userID, year, month)
	if err != nil {
		log.Error("Unable to execute query", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Transaction
		var amountStr string
		err := rows.Scan(&t.Timestamp, &t.Description, &amountStr)
		if err != nil {
			log.Error("Error scanning transaction", err)
			continue
		}
		t.Amount, _ = decimal.NewFromString(amountStr) // Handle error appropriately
		transactions = append(transactions, t)
	}
	return transactions, nil
}
