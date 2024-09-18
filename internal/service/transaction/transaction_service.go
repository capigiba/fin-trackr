package transaction

import (
	"fintrack/internal/domain/model"
	"fintrack/internal/repo/transaction"
)

type TransactionService struct {
	repo *transaction.TransactionRepository
}

func NewTransactionService(repo *transaction.TransactionRepository) *TransactionService {
	return &TransactionService{repo}
}

func (s *TransactionService) GetTransactions(userID uint64, year, month int) ([]model.Transaction, error) {
	return s.repo.Load(userID, year, month)
}
