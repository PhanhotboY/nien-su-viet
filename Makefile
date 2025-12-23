exec-%:
	docker compose exec -it $* bash

migrate-%:
	bun --cwd apps/$* prisma migrate dev

setup-%:
	bun --cwd apps/$* setup/index.ts

OPENAPI_DIR = openapi
OUTPUT_DIR = libs/nsv-interfaces

JSONS := $(wildcard $(OPENAPI_DIR)/*.json)
TS_OUTPUTS := $(patsubst $(OPENAPI_DIR)/%.json,$(OUTPUT_DIR)/%.ts,$(JSONS))

# Default target: generate all
gen-types: clean-types $(TS_OUTPUTS)

# Rule: convert one .json → one .ts
$(OUTPUT_DIR)/%.ts: $(OPENAPI_DIR)/%.json
	@echo "Generating types for $< → $@"
	# @bunx @aopture/openapi-down-convert --input $< --output $@
	@bunx -y openapi-typescript $< -o $@

# Clean generated files
clean-types:
	@rm -f $(OUTPUT_DIR)/*.ts

swagger2openapi:
	@echo "Converting swagger to openapi"
	@mkdir -p $(OPENAPI_DIR)
	@for file in $(wildcard swagger/*.json); do \
		base=$$(basename $$file .json); \
		echo "Processing $$file → $(OPENAPI_DIR)/$$base.json"; \
		swagger2openapi --outfile $(OPENAPI_DIR)/$$base.json $$file; \
	done

.PHONY: gen-types clean-types
