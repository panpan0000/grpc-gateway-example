#This is a golang GRPC Gateway example.


#开发指南 Developer Guide

refer to tutorial:
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/

Based on last grpc example (NO HTTP, Only GRPC) : https://github.com/panpan0000/grpc-example 

- 注意`*.proto` File 中,要加上HTTP->GRPC 的mapping
- 手动拷贝一些依赖库: myproto/google/ 
- Add `--grpc-gateway_out` as `protoc` parameter. to generate `*.pb.gw.go`
-  server的代码有较大的改动，主要是一个客户端连接grpc server后转成http服务



#使用方式：

1. 服务器端 `make run-server`
2. 客户端`curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": "Peter"}' `


#工具链版本

```
#go version
go version go1.17.5 darwin/amd64

# Proto toolchain versions:
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0 
	google.golang.org/protobuf v1.27.1 

```
