package service

import (
	"fintrack/internal/infra/env"
	"fintrack/internal/service/auth"
	"fintrack/internal/service/transaction"
	"fintrack/internal/service/user"

	"github.com/google/wire"
)

var SecretKey = env.EnvConfig.JWTSecret

// ProviderSetService is providers.
var ProviderSetService = wire.NewSet(
	user.NewUserService,
	transaction.NewTransactionService,
	auth.NewAuthService,
	wire.Value(SecretKey),
)
