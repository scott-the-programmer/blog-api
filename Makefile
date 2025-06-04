
.PHONY: tidy build clean test run fmt vet lint help

# Default target
all: tidy fmt vet test build

# Go module maintenance
tidy:
	@echo "Tidying go modules..."
	go mod tidy

# Build the application
build:
	@echo "Building application..."
	go build -o bin/blog-api ./main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Vet code
vet:
	@echo "Vetting code..."
	go vet ./...

# Run the application
run:
	@echo "Running application..."
	go run main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download

# Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2"; \
	fi

# Watch for changes and rebuild (requires air)
watch:
	@echo "Starting file watcher..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air not found. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Then run 'make watch' again"; \
	fi

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Run tidy, fmt, vet, test, and build"
	@echo "  tidy         - Tidy go modules"
	@echo "  build        - Build the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  run          - Run the application"
	@echo "  deps         - Install dependencies"
	@echo "  lint         - Run linter (requires golangci-lint)"
	@echo "  watch        - Watch for changes and rebuild (requires air)"
	@echo "  help         - Show this help message"
