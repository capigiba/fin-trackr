package handler

import (
	"fintrack/internal/handler/rest/v1/transaction"
	"fintrack/internal/handler/rest/v1/user"

	"github.com/google/wire"
)

// ProviderSetHandler is Handler providers.
var ProviderSetHandler = wire.NewSet(
	user.NewUserHandler,
	transaction.NewTransactionHandler,
)
