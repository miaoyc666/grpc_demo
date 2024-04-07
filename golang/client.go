package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "demo/api" // 替换为实际生成的 proto 文件的 Go 包路径

	"google.golang.org/grpc"
)

func main() {
	// 建立连接到 gRPC 服务
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStringServiceClient(conn)

	// 调用服务
	name := "hello world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ToUpper(ctx, &pb.StringMessage{Value: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetValue())
}
