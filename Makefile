.PHONY: build run clean test help install-deps init-submodules update-submodules

# Build the application
build:
	@echo "Building the application..."
	go build -o poc-server .

# Run the application
run: build
	@echo "Starting the server on :8080..."
	./poc-server

# Run without building (using go run)
dev:
	@echo "Running in development mode..."
	go run main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f poc-server
	go clean

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Initialize git submodules
init-submodules:
	@echo "Initializing git submodules..."
	git submodule update --init --recursive

# Update git submodules to latest commit
update-submodules:
	@echo "Updating git submodules..."
	git submodule update --remote --merge

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Skipping..."; \
	fi

# Show help
help:
	@echo "Available targets:"
	@echo "  build             - Build the application"
	@echo "  run               - Build and run the application"
	@echo "  dev               - Run in development mode (without building binary)"
	@echo "  clean             - Clean build artifacts"
	@echo "  install-deps      - Install Go dependencies"
	@echo "  init-submodules   - Initialize git submodules"
	@echo "  update-submodules - Update git submodules to latest"
	@echo "  test              - Run tests"
	@echo "  fmt               - Format code"
	@echo "  lint              - Lint code (requires golangci-lint)"
	@echo "  help              - Show this help message"

# Default target
.DEFAULT_GOAL := help
