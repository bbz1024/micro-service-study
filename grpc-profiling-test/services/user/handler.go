package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-profiling-test/rpc/goods"
	"grpc-profiling-test/rpc/user"
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

func Fob(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fob(n-1) + Fob(n-2)
}
func MakeArea() {
	c := make([]byte, 1024*8+1024)
	fmt.Println(c)
}
func (u *User) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	//time.Sleep(time.Second * 10)
	//fb := Fob(request.Id)
	//MakeArea()

	//获取商品信息
	goodsRes, err := goodClient.GetGoodsById(ctx, &goods.GoodsRequest{
		Id: 10,
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
