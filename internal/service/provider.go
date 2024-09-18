package service

import (
	"fintrack/internal/infra/env"

	"github.com/google/wire"
)

var SecretKey = env.EnvConfig.JWTSecret

// ProviderSetService is providers.
var ProviderSetService = wire.NewSet()
