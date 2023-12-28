package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-loadbalance-test/pkg/discovery"
	"grpc-loadbalance-test/rpc/user"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
)

var (
	ports = []string{":10001", ":10002", ":10003"}
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	var s []*grpc.Server
	for _, p := range ports {
		go func(port string) {
			rpcServer := grpc.NewServer(
				grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
					fmt.Println("负载均衡", port)
					resp, err = handler(ctx, req)
					return
				}),
			)
			//注册服务
			user.RegisterUserServiceServer(rpcServer, &User{})
			listen, err := net.Listen("tcp", port)
			if err != nil {
				panic(err)
			}
			iPort, _ := strconv.Atoi(port[1:])
			err = discovery.Consul.Register(
				context.Background(), discovery.Service{
					Name: "user",
					Host: "127.0.0.1",
					Port: iPort,
				},
			)
			if err != nil {
				panic(err)
				return
			}
			log.Println("user 服务启动成功 http://127.0.0.1" + port)
			s = append(s, rpcServer)
			rpcServer.Serve(listen)
		}(p)
	}
	stop := <-c
	for i, server := range s {
		server.GracefulStop()
		server.Stop()
		log.Println("user 服务停止成功", ports[i])
	}
	fmt.Println(stop)
}
