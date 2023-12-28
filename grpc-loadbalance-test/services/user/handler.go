package main

import (
	"context"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"grpc-loadbalance-test/rpc/user"
)

var (
	Users = map[int64]*user.UserResponse{
		1: {
			Name: "bbz",
			Age:  21,
		},
		2: {
			Name: "bbz",
			Age:  21,
		},
		3: {
			Name: "bbz",
			Age:  21,
		},
	}
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {

	return Users[request.Id], nil
}
