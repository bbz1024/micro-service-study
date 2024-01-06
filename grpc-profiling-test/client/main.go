package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-profiling-test/rpc/user"
	"log"
)

const ADDR = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	for i := 0; i < 500; i++ {
		go func() {
			fmt.Println(i)
			res, err := client.GetUser(context.Background(), &user.UserRequest{Id: int64(10)})
			if err != nil {
				fmt.Println("err:", err)
				return
			} else {
				fmt.Println("user:", res.Name, res.Age)
				fmt.Println("goods:", res.Goods.Name, res.Goods.Price)
				fmt.Println("====================================")
			}
		}()
	}
	select {}

}
