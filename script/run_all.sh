#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Variables
CMD_DIR="cmd/server"
APP_NAME="fintrackr"

# Function to display a message in green
log_info() {
    echo -e "${GREEN}$1${NC}"
}

# Function to display a message in red
log_error() {
    echo -e "${RED}$1${NC}"
}

# Function to check the last command's exit status and exit on error
check_error() {
    if [ $? -ne 0 ]; then
        log_error "Error encountered. Exiting script."
        exit 1
    fi
}

# Step 1: Install dependencies
log_info "Step 1: Installing Go module dependencies..."
go install github.com/google/wire/cmd/wire@latest
check_error

go install github.com/swaggo/swag/cmd/swag@latest
check_error

go mod tidy
check_error

go mod vendor
check_error

# Step 2: Generate Wire dependencies
log_info "Step 2: Generating Wire dependencies..."
cd $CMD_DIR || { log_error "Failed to change directory to $CMD_DIR"; exit 1; }
wire
check_error
cd - || exit 1 # Go back to the root directory

# Step 3: Build the application
log_info "Step 3: Building the application..."
# cd $CMD_DIR && go build -o $APP_NAME
cd $CMD_DIR && go build -o $APP_NAME -ldflags "-s -w"

cd -

# Step 4: Generate Swagger documentation
log_info "Step 4: Generating Swagger documentation..."
swag init -g $CMD_DIR/main.go -o ./docs

# Step 5: Run the built application
log_info "Step 6: Running the built application..."
cd $CMD_DIR && ./$APP_NAME
