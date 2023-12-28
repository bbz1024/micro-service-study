package main

import (
	"context"
	"grpc-recovery-test/rpc/user"
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {

	panic("implement me")
}
