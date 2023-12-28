package main

import (
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"grpc-trace-test/pkg/tracing"
	"grpc-trace-test/rpc/goods"
	"log"
	"net"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: main
 * @Date:
 * @Desc: ...
 *
 */
const (
	port = ":50052"
)

func main() {
	tracer, closer := tracing.Init("GoodsService")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			//otelgrpc.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)
	//注册服务
	goods.RegisterGoodsServiceServer(rpcServer, &Goods{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("goods 服务启动成功 http://127.0.0.1:50052")
	rpcServer.Serve(listen)
}
