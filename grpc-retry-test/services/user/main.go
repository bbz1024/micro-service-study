package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-retry-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

func main() {
	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
