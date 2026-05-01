.PHONY: help build run test clean docker-build docker-run

help:
	@echo "Available commands:"
	@echo "  make build         - Build the Go application"
	@echo "  make run           - Run the application"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"

env:
	@if [ ! -f .env ]; then cp .env.example .env; fi

build: env
	go build -o bin/cash-flow-forecast .

run: env
	go run main.go

dev: env
	go run main.go

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	coverage: coverage

clean:
	rm -rf bin/
	rm -f coverage.out

deps:
	go mod download
	deps-tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-build:
	docker build -t cash-flow-forecast:latest .

docker-run: docker-build
	docker run --rm -p 8080:8080 --env-file .env cash-flow-forecast:latest

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down
