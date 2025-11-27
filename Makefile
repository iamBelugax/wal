# ANSI Color Codes
CYAN := \033[36m
RESET := \033[0m
GREEN := \033[32m
YELLOW := \033[33m

MODULE_PATH := github.com/iamBelugax/wal

PROTO_DIR := internal/encoding/proto
PROTO_OUT_DIR := internal/encoding/proto/__gen__

tidy:
	@echo "$(CYAN) Tidying Go modules...$(RESET)"
	@go mod tidy
	@echo "$(GREEN) Go modules tidied.$(RESET)"

deps:
	@echo "$(CYAN) Downloading Go modules...$(RESET)"
	@go mod download
	@go mod verify
	@echo "$(GREEN) Go modules downloaded.$(RESET)"

fmt:
	@echo "$(CYAN) Formatting Go code...$(RESET)"
	@go fmt ./...
	@echo "$(GREEN) Formatting complete.$(RESET)"

gen-pb: clean-proto-gen
	@echo "$(CYAN) Generating Protocol Buffer and GRPC Go code...$(RESET)"
	@mkdir -p $(PROTO_OUT_DIR)
	@protoc \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=module=$(MODULE_PATH) \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go-grpc_opt=module=$(MODULE_PATH) \
		--proto_path=$(PROTO_DIR) \
		$(PROTO_DIR)/wal.proto
	@echo "$(GREEN) Protocol Buffer and GRPC generation complete$(RESET)"

clean-proto-gen:
	@echo "$(YELLOW) Cleaning previous Protocol Buffer and GRPC generated files...$(RESET)"
	@rm -rf $(PROTO_OUT_DIR)
	@echo "$(GREEN) Cleanup complete$(RESET)"