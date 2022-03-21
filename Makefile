SHELL := /bin/bash -o pipefail
gen-proto:
	protoc -I ./myproto \
   --go_out ./myproto --go_opt paths=source_relative \
   --go-grpc_out ./myproto --go-grpc_opt paths=source_relative \
   ./myproto/*.proto

run-server: gen-proto
	@go mod tidy
	go run server.go

.PHONY: gen-proto run-server

