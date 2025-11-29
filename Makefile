.PHONY: help tidy deps fmt gen-pb clean-proto-gen build clean run lint vet install

# Color definitions
BOLD := \033[1m
RESET := \033[0m
CYAN := \033[0;36m
GREEN := \033[0;32m
WHITE := \033[0;37m

# Configuration
BINARY_NAME := wal
CMD_DIR := cmd/wal
PROTO_DIR := internal/encoding/proto
MODULE_PATH := github.com/iamBelugax/wal
PROTO_OUT_DIR := internal/encoding/proto/__gen__

help: ## Display this help message
	@echo "$(BOLD)$(CYAN)╔════════════════════════════════════════════════════════════╗$(RESET)"
	@echo "$(BOLD)$(CYAN)║          WAL (Write-Ahead Log) - Build System              ║$(RESET)"
	@echo "$(BOLD)$(CYAN)╚════════════════════════════════════════════════════════════╝$(RESET)"
	@echo "$(BOLD)$(WHITE)Available targets:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(BOLD)$(GREEN)%-20s$(RESET) %s\n", $$1, $$2}'

tidy: ## Tidy Go modules
	@go mod tidy

deps: ## Download and verify dependencies
	@go mod download
	@go mod verify

fmt: ## Format Go code
	@go fmt ./...

gen-pb: clean-proto-gen ## Generate Protocol Buffer code
	@mkdir -p $(PROTO_OUT_DIR)
	@protoc \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=module=$(MODULE_PATH) \
		--proto_path=$(PROTO_DIR) \
		$(PROTO_DIR)/wal.proto

clean-proto-gen: ## Clean generated protobuf files
	@rm -rf $(PROTO_OUT_DIR)

build: gen-pb ## Build the WAL binary
	@go build -o bin/$(BINARY_NAME) ./$(CMD_DIR)

lint: ## Run golangci lint
	@golangci-lint run ./...

vet: ## Run go vet
	@go vet ./...

clean: clean-proto-gen ## Clean build artifacts and generated files
	@rm -rf bin/

run: build ## Build and run the WAL application
	@./bin/$(BINARY_NAME)

all: clean deps fmt gen-pb run ## Run all build steps