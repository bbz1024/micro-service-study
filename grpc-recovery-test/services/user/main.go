package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"grpc-recovery-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

func main() {

	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(
				func(ctx context.Context, p any) (err error) {
					fmt.Println("panic 恢复", p)
					return errors.New("server panic")
				})),
		),
	)
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
