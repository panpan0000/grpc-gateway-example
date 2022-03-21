SHELL := /bin/bash -o pipefail
gen-proto:
	pushd myproto && \
	protoc --go_out=. --go-grpc_out=. *.proto ; \
	popd

run-server: gen-proto
	@go mod tidy
	go run server.go

run-client: gen-proto
	@go mod tidy
	go run client.go


.PHONY: gen-proto run-server run-client 

