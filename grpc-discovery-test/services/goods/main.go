package main

import (
	"context"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-discovery-test/pkg/discovery"
	"grpc-discovery-test/rpc/goods"
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

	rpcServer := grpc.NewServer()
	//注册服务
	goods.RegisterGoodsServiceServer(rpcServer, &Goods{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	err = discovery.Consul.Register(
		context.Background(), discovery.Service{
			Name: "goods",
			Host: "127.0.0.1",
			Port: 50052,
		},
	)
	if err != nil {
		panic(err)

	}

	log.Println("goods 服务启动成功 http://127.0.0.1:50052")
	rpcServer.Serve(listen)
}
