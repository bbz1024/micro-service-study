package main

import (
	"context"
	"errors"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"google.golang.org/grpc"
	"grpc-breake-test/rpc/goods"
	"grpc-breake-test/rpc/user"
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
	e, b := sentinel.Entry("get_user")
	if b != nil {
		return &user.UserResponse{
			Name: "not found user",
			Age:  0,
		}, b
	}
	defer e.Exit()
	if true {
		sentinel.TraceError(e, errors.New("biz error"))
		return nil, errors.New("test sentinel strategy error")
	}
	//获取商品信息
	goodsRes, err := goodClient.GetGoodsById(ctx, &goods.GoodsRequest{
		Id: int64(rand.Int31() % 10),
	})
	if err != nil {
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
func init() {

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	goodClient = goods.NewGoodsServiceClient(conn)

}
