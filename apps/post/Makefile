.PHONY dev:
	@echo "Running post service..."
	@air

.PHONY clean:
	@echo "Cleaning post service..."
	@rm -rf bin

.PHONY build: clean
	@echo "Building post service..."
	@go build -o bin/server cmd/server/main.go
