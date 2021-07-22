package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	service "grpc/service"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := service.NewCalClient(conn)
	// 调用服务端函数
	r, err := c.Multiply(context.Background(), &service.MultiplyRequest{
		Number1: 100,
		Number2: 300,
	})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Printf("调用成功: %d\n", r.Number)
	// 调用服务端函数
	r2, err2 := c.Add(context.Background(), &service.AddRequest{
		Number1: 100,
		Number2: 300,
	})
	if err2 != nil {
		fmt.Printf("调用服务端代码失败: %s", err2)
		return
	}
	fmt.Printf("调用成功: %d\n", r2.Number)
}
