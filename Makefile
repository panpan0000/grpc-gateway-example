SHELL := /bin/bash -o pipefail
gen-proto:
	protoc -I ./myproto \
   --go_out ./myproto  \
   --go-grpc_out ./myproto  \
   --grpc-gateway_out ./myproto \
   ./myproto/*.proto

run-server: gen-proto
	@go mod tidy
	go run server.go

.PHONY: gen-proto run-server

