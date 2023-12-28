package main

import (
	"fmt"
	"google.golang.org/grpc"
	"micro-service/grpc-pool-test/rpc/user"
	"net"
)

func main() {


	rpcServer := grpc.NewServer()

	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer.Serve(listen)
}
