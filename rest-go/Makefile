SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    build-server             Build the server executable"
	@echo "    run-server               Run the server"
	@echo ""

.PHONY: build-server
build-server:
	go build -o bin/soa-server ./src/boot.go 
	cp ./src/resources/pubs.json ./bin/resources

.PHONY: run-server
run-server:
	(cd ./bin; ./soa-server)
