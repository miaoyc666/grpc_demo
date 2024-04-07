package main

import (
	"context"
	"log"
	"net"
	"strings"

	pb "demo/api" // 替换为实际生成的 proto 文件的 Go 包路径
	"google.golang.org/grpc"
)

// stringServiceServer 是 pb.StringServiceServer 的实现
type stringServiceServer struct {
	pb.UnimplementedStringServiceServer
}

// ToUpper 实现了 StringService 服务中定义的 ToUpper 方法
func (s *stringServiceServer) ToUpper(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: strings.ToUpper(in.Value)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &stringServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
