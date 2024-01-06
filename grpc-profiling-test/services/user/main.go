package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-profiling-test/pkg/profiling"
	"grpc-profiling-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

// Fob 计算一个fob

func main() {
	profiling.InitProfiling("user9")

	rpcServer := grpc.NewServer()
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//开启profiling
	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
