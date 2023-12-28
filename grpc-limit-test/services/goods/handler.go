package main

import (
	"context"
	"grpc-limit-test/rpc/goods"
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
	//time.Sleep(time.Millisecond * 400)
	if g, ok := goodsMap[int32(request.Id)]; ok {
		return g, nil
	}
	return &goods.GoodsResponse{Name: "not found goods", Price: 0}, nil

}
