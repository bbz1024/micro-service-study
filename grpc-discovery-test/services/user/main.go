package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-discovery-test/pkg/discovery"
	"grpc-discovery-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

func main() {

	rpcServer := grpc.NewServer()
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = discovery.Consul.Register(
		context.Background(), discovery.Service{
			Name: "user",
			Host: "127.0.0.1",
			Port: 50051,
		},
	)
	if err != nil {
		panic(err)
		return
	}

	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
