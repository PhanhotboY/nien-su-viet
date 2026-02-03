APP_NAME = server
CMS_SERVER_ENV ?= development

generate:
	go run cmd/gen-openapi/main.go

dev:
	nodemon --config ./nodemon.json cmd/$(APP_NAME)/main.go

clean:
	rm -rf bin/$(APP_NAME)

build: clean
	go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

run: build
	CMS_SERVER_ENV=production ./bin/$(APP_NAME)

OPENAPI_DIR = ../../openapi
OUTPUT_DIR = pkg/generated

JSONS := $(wildcard $(OPENAPI_DIR)/*.json)
GO_OUTPUTS := $(patsubst $(OPENAPI_DIR)/%.json,$(OUTPUT_DIR)/%.go,$(JSONS))

gen-types: clean-types $(GO_OUTPUTS)

# Rule: convert one .json → one .go
$(OUTPUT_DIR)/%.go: $(OPENAPI_DIR)/%.json
	@echo "Generating types for $< → $@"
	@$(shell go env GOPATH)/bin/oapi-codegen -package api $< > $@


# Clean generated files
clean-types:
	@rm -f $(OUTPUT_DIR)/*.go


.PHONY: dev clean build run
