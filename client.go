package main

import (
    "context"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "myproto"
)

const (
    address     = "localhost:50001"
    defaultName = "Peter's Client"
)

func main() {
    //建立链接
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewMyEchoClient(conn)

    // Contact the server and print out its response.
    name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    // 1秒的上下文超时
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.Foo(ctx, &pb.MyGrpcRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Received from Server(%s) : %s",address,  r.Message)
}
