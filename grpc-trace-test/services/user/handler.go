package main

import (
	"context"
	"grpc-trace-test/pkg/conn"
	"grpc-trace-test/pkg/tracing"
	"grpc-trace-test/rpc/goods"
	"grpc-trace-test/rpc/user"
	"math/rand"
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
var (
	goodClient goods.GoodsServiceClient
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	//1. //创建span
	span := tracing.StartSpan(ctx, "GetUser")
	tracing.RecordWithIP(span, "8.8.8.8")
	defer span.Finish()
	New(ctx)
	//获取商品信息
	goodsRes, err := goodClient.GetGoodsById(ctx, &goods.GoodsRequest{
		Id: int64(rand.Int31() % 10),
	})
	if err != nil {
		tracing.RecordError(span, err)
		res := new(user.UserResponse)
		res.Name = "not found user"
		res.Age = 0

		res.Goods = &goods.GoodsResponse{
			Name:  "not found goods",
			Price: 0,
		}
		return res, err
	}
	if obj, ok := Users[request.Id]; ok {
		obj.Goods = goodsRes
		return obj, nil
	}
	return &user.UserResponse{
		Goods: goodsRes,
		Name:  "not found user",
		Age:   0,
	}, nil
}
func New(ctx context.Context) {
	clientConn, err := conn.Conn(ctx, "127.0.0.1:50052")
	if err != nil {
		panic(err)
	}
	goodClient = goods.NewGoodsServiceClient(clientConn)
}
