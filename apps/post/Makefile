.PHONY: dev
dev:
	@echo "Running post service..."
	@air

.PHONY: clean
clean:
	@echo "Cleaning post service..."
	@rm -rf bin

.PHONY: build
build: clean
	@echo "Building post service..."
	@go build -o bin/server cmd/server/main.go

.PHONY: test
test:
	@echo "Running tests for post service..."
	@POST_SERVER_ENV=test go test -v ./... -cover -coverprofile=coverage.out -coverpkg=./internal/...
	@make open-coverage

test-%:
	@echo "Running tests for post service with pattern: $*..."
	@POST_SERVER_ENV=test go test -v -run $* ./... -coverprofile=coverage.out -coverpkg=./internal/...
	@make create-coverage

.PHONY: create-coverage
create-coverage:
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: open-coverage
open-coverage:
	@echo "Opening coverage report for post service..."
	@make create-coverage
	@xdg-open coverage.html > /dev/null 2>&1 &
