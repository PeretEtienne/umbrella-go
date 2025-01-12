.PHONY: run build clean test

# Variables
APP_NAME = my-gin-app
DOCKER_COMPOSE = docker compose
AIR_CMD = air
GO_CMD = go
DB_CONTAINER = postgres

# Run the application with live reload
run:
	@$(DOCKER_COMPOSE) up --build -d

# Build the application binary
build:
	@$(GO_CMD) build -o bin/$(APP_NAME)

# Clean the build artifacts
clean:
	@rm -rf bin/

# Run tests
test:
	@$(GO_CMD) test ./... -v

