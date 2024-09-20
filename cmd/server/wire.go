// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	handler "fintrack/internal/handler/rest/v1"
	"fintrack/internal/infra/db"
	"fintrack/internal/pkg/middleware"
	"fintrack/internal/repo"
	"fintrack/internal/router"
	"fintrack/internal/service"

	"github.com/google/wire"
)

// InitializeApp wires dependencies using Wire and should not have manual code.
func InitializeApp() (*router.AppRouter, error) {
	wire.Build(
		db.ProviderSetDB,
		repo.ProviderSetRepository,
		service.ProviderSetService,
		handler.ProviderSetHandler,
		middleware.ProviderSetMiddleware,
		router.ProviderSetRouter,
	)
	return nil, nil
}
