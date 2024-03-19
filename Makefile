# Variables
PROTO_DIR := internal/api/greeterService
PROTO_FILE := $(PROTO_DIR)/greeterService.proto
PROTO_GO_OUT := $(PROTO_DIR)

# Targets
.PHONY: all
all: generate

.PHONY: generate
generate:
	@echo "Generating Go code from protobuf..."
	protoc --go_out=$(PROTO_GO_OUT) --go-grpc_out=$(PROTO_GO_OUT) $(PROTO_FILE)
	@echo "Go code generation complete"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -f $(PROTO_GO_OUT)/*.pb.go
	@echo "Cleaned generated files"
	@echo "Cleaning cert files..."
	rm -f certs/client/*.pem certs/server/*.pem certs/CA/*.pem certs/client/*.p12 certs/server/*.p12
	@echo "Cleaned cert files"


.PHONY: help
help:
	@echo "Available targets:"
	@echo "  generate  : Generate Go code from protobuf file"
	@echo "  clean     : Clean generated files"
	@echo "  help      : Show this help message"

.DEFAULT_GOAL := help
