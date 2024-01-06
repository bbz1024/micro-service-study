package main

import (
	"context"
	"grpc-prometheus-test/rpc/user"
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {

	return &user.UserResponse{
		Name: "not found user",
		Age:  0,
	}, nil
}
