package transaction

import (
	"fintrack/internal/domain/entity"
	"fintrack/internal/domain/model"
	"fintrack/internal/repo/transaction"
)

type TransactionService struct {
	repo *transaction.TransactionRepository
}

func NewTransactionService(r *transaction.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) SaveTransaction(t entity.Transaction) error {
	return s.repo.Save(t)
}

func (s *TransactionService) GetTransactions(userID uint64, year, month int) ([]model.Transaction, error) {
	return s.repo.Load(userID, year, month)
}
