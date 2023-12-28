package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-breake-test/rpc/user"
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
	for i := 0; i < 10; i++ {
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
