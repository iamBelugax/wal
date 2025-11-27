PROTO_DIR := internal/encoding/proto
MODULE_PATH := github.com/iamBelugax/wal
PROTO_OUT_DIR := internal/encoding/proto/__gen__

tidy:
	@go mod tidy

deps:
	@go mod download
	@go mod verify

fmt:
	@go fmt ./...

gen-pb: clean-proto-gen
	@mkdir -p $(PROTO_OUT_DIR)
	@protoc \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=module=$(MODULE_PATH) \
		--proto_path=$(PROTO_DIR) \
		$(PROTO_DIR)/wal.proto

clean-proto-gen:
	@rm -rf $(PROTO_OUT_DIR)