package main

import (
	"context"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-loadbalance-test/pkg/discovery"
	"grpc-loadbalance-test/rpc/goods"
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
	port = ":20000"
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
			Port: 20000,
		},
	)
	if err != nil {
		panic(err)

	}

	log.Println("goods 服务启动成功 http://127.0.0.1:20000")
	rpcServer.Serve(listen)
}
