package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	conn2 "grpc-trace-test/pkg/conn"
	"grpc-trace-test/pkg/tracing"
	"grpc-trace-test/rpc/user"
	"log"
)

const ADDR = "127.0.0.1:50051"

func main() {
	tracer, closer := tracing.Init("Client")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	conn, err := conn2.Conn(context.Background(), ADDR)
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	for i := 0; i < 1; i++ {
		res, err := client.GetUser(context.Background(), &user.UserRequest{Id: int64(i)})
		if err != nil {
			fmt.Println("err:", err)
			continue
		} else {
			fmt.Println("user:", res.Name, res.Age)
			fmt.Println("goods:", res.Goods.Name, res.Goods.Price)
			fmt.Println("====================================")
		}
	}

}
