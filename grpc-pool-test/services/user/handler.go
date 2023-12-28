package main

import (
	"context"
	"log"
	"micro-service/grpc-pool-test/rpc/user"
)

//grpc-pool使用

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Println("get user", request.Id)
	return &user.UserResponse{Name: "bbz", Age: 21}, nil

}
