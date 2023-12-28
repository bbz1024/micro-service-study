package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 使用consul作为grpc的服务发现
	"google.golang.org/grpc"
	"grpc-discovery-test/pkg/discovery"
	"grpc-discovery-test/rpc/goods"
	"grpc-discovery-test/rpc/user"
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
	addr, err := discovery.Consul.GetService(context.Background(), "goods")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	goodClient = goods.NewGoodsServiceClient(conn)

}
