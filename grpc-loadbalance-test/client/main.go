package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-loadbalance-test/pkg/discovery"
	"grpc-loadbalance-test/rpc/user"
	"log"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for i := 0; i < 10; i++ {
		go func() {
			addr, err2 := discovery.Consul.GetService(context.Background(), "user")
			if err2 != nil {
				return
			}
			conn, err := grpc.Dial(
				addr, grpc.WithInsecure(),
				//基于轮询算法
				grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
			)

			if err != nil {
				log.Fatalf("failed to get conn: %v", err)
			}
			defer conn.Close()
			client := user.NewUserServiceClient(conn)
			_, err = client.GetUser(context.Background(), &user.UserRequest{
				Id: 1,
			})

			if err != nil {
				log.Fatalf("failed to get user: %v", err)
			}

		}()
	}
	stop := <-c
	fmt.Println(stop)
}
