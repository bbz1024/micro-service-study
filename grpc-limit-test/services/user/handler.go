package main

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-limit-test/rpc/goods"
	"grpc-limit-test/rpc/user"
	"log"
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
	//基于sentinel的限流

	e, err := sentinel.Entry("get_user")
	if err != nil {
		log.Printf("Failed to enter sentinel: %s", err)
		return nil, status.Errorf(codes.ResourceExhausted, "sentinel limit exceeded")
	}
	defer e.Exit(base.WithError(err))
	//获取商品信息

	goodsRes, err2 := goodClient.GetGoodsById(ctx, &goods.GoodsRequest{
		Id: int64(rand.Int31() % 10),
	})
	if err2 != nil {
		res := new(user.UserResponse)
		res.Name = "not found user"
		res.Age = 0
		res.Goods = &goods.GoodsResponse{
			Name:  "not found goods",
			Price: 0,
		}
		return res, err2
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
func init() {

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	goodClient = goods.NewGoodsServiceClient(conn)

}
