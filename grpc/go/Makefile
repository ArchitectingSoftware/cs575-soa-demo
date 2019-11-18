SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    build-client             Build the client executable"
	@echo "    build-server             Build the server executable"
	@echo ""
	@echo "    generate               	Generate Protocol Buffers"
	@echo ""
	@echo "    run-server               Run the server"
	@echo "    run-client               Run the client"
	@echo ""

.PHONY: build-server
build-server:
	go build -o bin/server ./src/server
	cp ./src/server/pubs.json ./bin

.PHONY: run-server
run-server:
	(cd ./bin; ./server)

.PHONY: generate
generate:
	protoc -I src/pubs-proto/ src/pubs-proto/pubs.proto --go_out=plugins=grpc:src/pubs-proto

.PHONY: build-client
build-client:
	go build -o bin/client ./src/client 

.PHONY: run-client
run-client:
	(cd ./bin; ./client)
