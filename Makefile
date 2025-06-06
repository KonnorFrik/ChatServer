LOCAL_BIN=$(CURDIR)/bin
# Path for store all proto files
DEFAULT_API_PATH=api/proto
# Path for store generated output
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

# api/proto/
# 	serviceA/
# 		v1/
# 			file.proto
# 		v2/
# 			file.proto
# 	serviceB/
# 		v1/
# 			file.proto
# 		v2/
# 			file.proto
# name = 'serviceA' or 'serviceB' etc.
# version = '1' or '2' etc
# ver = shortened name for 'varsion'
version=$(ver)
# Directory with proto files for one service
PROTO_INPUT_DIR=$(DEFAULT_API_PATH)/$(name)/v$(version)
# Concrete input proto file
PROTO_INPUT_FILE=$(PROTO_INPUT_DIR)/$(name).proto

# Directory for output generated files
PROTO_OUTPUT_DIR=$(DEFAULT_PKG_PATH)/$(name)/v$(version)

# All proto files in single input dir for compilation
PROTO_INPUT_FILES=$(wildcard $(PROTO_INPUT_DIR)/*.proto)

# generate - generate a gRPC stubs from given $(name)/v$(version) folder
generate: check_name
	mkdir -p $(PROTO_OUTPUT_DIR)
	protoc -I $(PROTO_INPUT_DIR) \
		--go_out=$(PROTO_OUTPUT_DIR) --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=$(PROTO_OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		$(PROTO_INPUT_FILES) 

# check_name - check is vars 'name' and 'version' exists and is 'PROTO_INPUT_DIR' and 'PROTO_INPUT_FILE' exists before generate
check_name:
	@if [ -z "$(name)" ]; then echo "[!] Please provide a name for create"; exit 1 ; fi
	@if [ -z "$(version)" ]; then echo "[!] Please provide a version number for create"; exit 1 ; fi
	@if [ ! -d "$(PROTO_INPUT_DIR)" ]; then echo "[!] Dir $(PROTO_INPUT_DIR) is not exist"; exit 1 ; fi
	@if [ ! -e "$(PROTO_INPUT_FILE)" ]; then echo "[!] File '$(PROTO_INPUT_FILE)' is not exist"; exit 1; fi

# new - create new dir and first '.proto' file
new:
	@if [ -z "$(name)" ]; then echo "[!] Please provide a name for create"; exit 1 ; fi
	@if [ -z "$(version)" ]; then echo "[!] Please provide a version number for create"; exit 1 ; fi
	mkdir -p $(PROTO_INPUT_DIR)
	touch $(PROTO_INPUT_FILE)

# TODO: '.proto' template may be added in 'new' command

help:
	@echo "'make new name=<string> version=<int>'      : For create a new version of a gRPC proto files"
	@echo "'make generate name=<string> version=<int>' : For generate a gRPC stub"

