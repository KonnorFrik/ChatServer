LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=chat_server
DEFAULT_API_PATH=api/proto
DEFAULT_PKG_PATH=pkg

_DEFAULT: help

install_deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

get_deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

build: build_server

build_server:
	go build -o server cmd/server/main.go

#trigger for anything. Even on _DEFAULT :c
# %: 
# 	make generate name=$@


PROTO_INPUT_DIR=$(DEFAULT_API_PATH)/$(name)
PROTO_INPUT_FILE=$(PROTO_INPUT_DIR)/$(name).proto

PROTO_OUTPUT_DIR=$(DEFAULT_PKG_PATH)/$(name)

generate: check_name
	mkdir -p $(PROTO_OUTPUT_DIR)
	protoc -I $(PROTO_INPUT_DIR) \
		--go_out=$(PROTO_OUTPUT_DIR) --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=$(PROTO_OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		$(PROTO_INPUT_FILE)

check_name:
	@if [ -z "$(name)" ]; then echo "[!] Please provide a name for compilation"; exit 1 ; fi
	@if [ ! -d "$(PROTO_INPUT_DIR)" ]; then echo "[!] Dir $(PROTO_INPUT_DIR) is not exist"; exit 1 ; fi
	@if [ ! -e "$(PROTO_INPUT_FILE)" ]; then echo "[!] File '$(PROTO_INPUT_FILE)' is not exist"; exit 1; fi

new:
	@if [ -z "$(name)" ]; then echo "[!] Please provide a name for compilation"; exit 1 ; fi
	mkdir -p $(PROTO_INPUT_DIR)
	touch $(PROTO_INPUT_FILE)

help:
	@echo "'make new name=chat_server_X_Y'     : For create a new version of a gRPC proto files"
	@echo "'make generate name=chat_server_X_Y': For generate a gRPC stub"
	@echo "     where X - major and Y - minor versions of .proto file"
