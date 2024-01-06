package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-dtm-test/rpc/orders"
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

const Addr = "192.168.94.1:10000"

func main() {

	rpcServer := grpc.NewServer()
	//注册服务
	orders.RegisterOrdersServiceServer(rpcServer, &Order{})
	listen, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("orders 服务启动成功 ", Addr)
	rpcServer.Serve(listen)
}
