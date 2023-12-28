package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-retry-test/rpc/goods"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: handler
 * @Date:
 * @Desc: ...
 *
 */

type Goods struct {
	goods.UnimplementedGoodsServiceServer
}

var goodsMap = map[int32]*goods.GoodsResponse{
	1: {Name: "苹果", Price: 10},
	2: {Name: "小脚", Price: 20},
	3: {Name: "橘子", Price: 30},
}

func (u *Goods) GetGoodsById(ctx context.Context, request *goods.GoodsRequest) (*goods.GoodsResponse, error) {
	return nil, status.Errorf(codes.AlreadyExists, "重试测试")

}
