package main

import (
	"context"
	"log"
	"net"
    "net/http"
    "google.golang.org/grpc/credentials/insecure"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
    pb "myproto"
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
    log.Println("starting grpc server on port ", port)
    //////////// GRPC Server用go routine来运行//////////////
    go func() {
		log.Fatalln(s.Serve(lis))
	}()
    //这里有一个grpc的客户端，来连接刚才go routine启动的grpc服务端
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0" + port,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
        log.Fatalln("Failed to dial server:", err)
	}
    // 启动一个HTTP服务
	gwmux := runtime.NewServeMux()
	// Register PB服务
	err = pb.RegisterMyEchoHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
