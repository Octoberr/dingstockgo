package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	service "grpc/service"
	"net"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *service.AddRequest) (*service.ResultReply, error) {
	result := in.Number1 + in.Number2
	fmt.Println("Call Add function")
	return &service.ResultReply{Number: result}, nil
}

func (s *server) Multiply(ctx context.Context, in *service.MultiplyRequest) (*service.ResultReply, error) {
	result := in.Number1 * in.Number2
	fmt.Println("Call Mutiply function")
	return &service.ResultReply{Number: result}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println("Start list at 50051")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	service.RegisterCalServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}

}
