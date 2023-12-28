package main

import (
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"grpc-trace-test/pkg/tracing"
	"grpc-trace-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

func main() {
	tracer, closer := tracing.Init("UserService")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			//otelgrpc.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
		),
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
