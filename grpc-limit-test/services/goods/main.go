package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-limit-test/rpc/goods"
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
		fmt.Println(err)
		panic(err)
	}
	log.Println("goods 服务启动成功 http://127.0.0.1:50052")
	rpcServer.Serve(listen)
}
