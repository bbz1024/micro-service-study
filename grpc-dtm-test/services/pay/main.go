package main

import (
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"google.golang.org/grpc"
	"grpc-dtm-test/rpc/pay"
	"log"

	"net"
)

const Addr = "192.168.94.1:10001"

func main() {

	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(dtmgimp.GrpcServerLog))
	rpcServer.RegisterService(&pay.PayService_ServiceDesc, &Pay{})

	listen, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("pay 服务启动成功 ", Addr)
	rpcServer.Serve(listen)
}
