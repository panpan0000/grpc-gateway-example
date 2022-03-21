package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
    pb "myproto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50001"
)

// server is used to implement myproto.MyEcho
type server struct {
	pb.UnimplementedMyEchoServer
}

// 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) Foo(ctx context.Context, in *pb.MyGrpcRequest) (*pb.MyGrpcReply, error) {
	return &pb.MyGrpcReply{Message: "Echoing Client name " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起GRPC服务
	pb.RegisterMyEchoServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
    log.Println("starting grpc server on port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
