package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"grpc-retry-test/rpc/goods"
	"grpc-retry-test/rpc/user"
	"math/rand"
	"time"
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

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(retry.UnaryClientInterceptor(
			retry.WithMax(3),

			retry.WithOnRetryCallback(func(ctx context.Context, attempt uint, err error) {
				fmt.Println("重试回调方法", attempt, err)
			}),
			//每次重试的超时时间为 1 秒
			retry.WithPerRetryTimeout(1*time.Second),
			//指定允许重试的 gRPC 错误码。
			retry.WithCodes(codes.AlreadyExists, codes.Unavailable),
			//定义退避策略，该函数在每次重试时调用，并返回下一次重试前要等待的时间。在这里，每次重试都等待 3 秒。
			retry.WithBackoff(func(ctx context.Context, attempt uint) time.Duration {

				return 3 * time.Second
			}),
		)),
	)
	if err != nil {
		panic(err)
	}
	goodClient = goods.NewGoodsServiceClient(conn)

}
