package main

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-dtm-test/db"
	"grpc-dtm-test/rpc/pay"
	"log"
)

type Pay struct {
	pay.UnimplementedPayServiceServer
}

func (p *Pay) Pay(ctx context.Context, request *pay.PayRequest) (*pay.PayResponse, error) {

	log.Println("pay.....")
	//获取用户余额
	key := fmt.Sprintf(db.UserKey, request.UserId)
	balance, err := db.RedisClient.HGet(ctx, key, "balance").Int64()
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("get balance error: %v", err)).Err()
	}
	if balance < request.Price {
		//回滚事务
		return nil, status.New(codes.Aborted, dtmcli.ResultFailure).Err()
	}
	//扣除余额
	if err := db.RedisClient.HIncrBy(ctx, key, "balance", -request.Price).Err(); err != nil {
		return nil, err
	}

	res := new(pay.PayResponse)
	log.Println("pay success")
	return res, nil

}

func (p *Pay) PayRevert(ctx context.Context, request *pay.PayRequest) (*pay.PayResponse, error) {
	log.Println("pay revert.....")
	//获取用户余额
	key := fmt.Sprintf(db.UserKey, request.UserId)
	//扣除余额
	if err := db.RedisClient.HIncrBy(ctx, key, "balance", request.Price).Err(); err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("get balance error: %v", err)).Err()
	}
	res := new(pay.PayResponse)
	log.Println("pay revert success")
	return res, nil
}
