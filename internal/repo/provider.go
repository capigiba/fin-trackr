package repo

import (
	"fintrack/internal/repo/transaction"
	"fintrack/internal/repo/user"

	"github.com/google/wire"
)

// ProviderSetRepository is peoviders.
var ProviderSetRepository = wire.NewSet(
	user.NewUserRepo,
	transaction.NewTransactionRepository,
)
