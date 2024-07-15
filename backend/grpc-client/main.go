package main

import (
    "context"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "grpc-server/proto"
)

const (
    address     = "localhost:50051"
    defaultName = "world"
)

func main() {
    // サーバーへの接続を設定
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    // サーバーに連絡し、レスポンスを出力
    name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // context.Background() を ctx に変更
    response, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", response.GetMessage())
}
