.PHONY: dev
dev:
	@echo "Running billing service..."
	@air

.PHONY: clean
clean:
	@echo "Cleaning billing service..."
	@rm -rf bin

.PHONY: build
build: clean
	@echo "Building billing service..."
	@go build -o bin/server cmd/server/main.go

.PHONY: test-
test-%:
	@echo "Running tests for billing service with pattern: $*..."
	@ENV_PREFIX=billing BILLING_SERVER_ENV=test go test -tags=$* -v ./... -cover -coverprofile=coverage.out -coverpkg=./internal/...
	@make create-coverage

.PHONY: create-coverage
create-coverage:
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: open-coverage
open-coverage:
	@echo "Opening coverage report for billing service..."
	@make create-coverage
	@xdg-open coverage.html > /dev/null 2>&1 &
