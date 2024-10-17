# Variables
CMD_DIR := cmd/server
MIGRATIONS_DIR := cmd/migration
OUTPUT_DIR := internal/wire_gen
APP_NAME := fintrackr
SCRIPT_DIR := script

# Default target
all: build

# Run the application
run:
	cd $(CMD_DIR) && go run .

migration:
	cd $(MIGRATIONS_DIR) && go run migrate.go

# Build the application
build:
	cd $(CMD_DIR) && go build -o $(APP_NAME)

# Generate wire dependencies
wire:
	cd $(CMD_DIR) && wire

# Clean the generated binaries
clean:
	rm -f $(CMD_DIR)/$(APP_NAME)

# Generate wire and then build
wire-build: wire build

# Swagger
swag:
	swag init -g $(CMD_DIR)/main.go -o ./docs

# Run all steps from the script
run-all:
	bash $(SCRIPT_DIR)/run_all.sh

# Help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make run         Run the application"
	@echo "  make swag		  Run the swagger"
	@echo "  make build       Build the application"
	@echo "  make wire        Generate dependencies with Wire"
	@echo "  make clean       Clean the generated binaries"
	@echo "  make wire-build  Generate Wire dependencies and build the application"
	@echo "  make help        Show this help message"