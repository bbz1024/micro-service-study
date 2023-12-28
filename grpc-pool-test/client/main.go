package main

import (
	"context"
	"github.com/shimingyah/pool"
	"log"
	"micro-service/grpc-pool-test/rpc/user"
)

func main() {
	//
	p, err := pool.New("127.0.0.1:50051", pool.DefaultOptions)
	if err != nil {
		log.Fatalf("failed to new pool: %v", err)
	}
	defer p.Close()
	conn, err := p.Get()
	for i := 0; i < 50; i++ {
		go func() {
			if err != nil {
				log.Fatalf("failed to get conn: %v", err)
			}
			client := user.NewUserServiceClient(conn.Value())
			res, err := client.GetUser(context.Background(), &user.UserRequest{Id: 2})
			if err != nil {
				log.Fatalf("failed to get user: %v", err)
			}
			log.Println(res)
		}()
	}
	defer conn.Close()

}
