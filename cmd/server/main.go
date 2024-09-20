package main

import (
	"fintrack/internal/infra/env"
	"fintrack/internal/infra/server/http"
	"fintrack/internal/infra/zap-logging/log"
	"fintrack/internal/infra/zap-logging/zap"
	"fintrack/internal/pkg/reason"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	// Name is the name of the project
	Name = "fin-trackr"
	// Version is the version of the project
	Version = "0.0.0"
	// log level
	logLevel string
	// log path
	logPath string
)

func main() {
	if env.EnvConfig == nil {
		fmt.Errorf("EnvConfig not loaded")
	}

	logLevel = env.EnvConfig.LogLevel
	logPath = env.EnvConfig.LogPath

	logDir := filepath.Dir(logPath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create log directory: %v", err)
		}
	}

	// Initialize logging
	log.SetLogger(zap.NewLogger(log.ParseLevel(logLevel), zap.WithName(Name), zap.WithPath(logPath), zap.WithCallerFullPath()))

	appRouter, err := InitializeApp()
	if err != nil {
		log.Error(reason.FailedToInitializeRouter, err)
		os.Exit(1)
	}

	// Create a new Gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	appRouter.RegisterUserRoutes(api)
	appRouter.RegisterTransactionRoutes(api)
	appRouter.RegisterSwaggerRoutes(r.Group("/"))

	// Create the http server
	addr := ":" + env.EnvConfig.ServerPort
	server := http.NewServer(r, addr)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Info(reason.ServerShutdown.Message())
		if err := server.Shutdown(); err != nil {
			log.Warn((reason.ServerForcedShutdown), err)
		}
		log.Info(reason.ServerExiting.Message())
	}()

	// Start the server
	if err := server.Start(); err != nil {
		log.Error(reason.ServerStartFailed.Message(), err)
	}
}
