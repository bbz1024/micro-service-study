package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-dtm-workflow-test/db"
	"grpc-dtm-workflow-test/rpc/orders"
	"log"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: handler
 * @Date:
 * @Desc: ...
 *
 */

type Order struct {
	orders.UnimplementedOrdersServiceServer
}

func (o *Order) Create(ctx context.Context, request *orders.OrdersRequest) (*orders.OrdersResponse, error) {
	log.Println("order.....")
	key := fmt.Sprintf(db.OrderKey, request.Id, request.UserId)
	data, err := json.Marshal(request)
	if err != nil {
		return nil, status.New(codes.Aborted, "FAILURE").Err()
	}

	if err := db.RedisClient.Set(ctx, key, data, 0).Err(); err != nil {

		return nil, status.New(codes.Aborted, "FAILURE").Err()
	}
	res := new(orders.OrdersResponse)
	log.Println("order success")
	return res, nil
}

func (o *Order) CreateRevert(ctx context.Context, request *orders.OrdersRequest) (*orders.OrdersResponse, error) {
	log.Println("order revert.....")
	key := fmt.Sprintf(db.OrderKey, request.Id, request.UserId)
	if err := db.RedisClient.Del(ctx, key).Err(); err != nil {
		return nil, status.New(codes.Aborted, "FAILURE").Err()
	}
	log.Println("order revert success")
	return &orders.OrdersResponse{}, nil
}
