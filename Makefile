.PHONY: test clean deps fmt lint mocks docs build run

tidy:
	go mod tidy

# Build the application
build:
	go build -o bin/app main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Generate mocks
mocks:
	mockgen -source=internal/domain/interfaces.go -destination=internal/mocks/interfaces_mock.go

# Generate API docs
docs:
	swag init -g internal/handler/*.go -o api/docs