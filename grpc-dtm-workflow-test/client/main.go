package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/workflow"
	"github.com/lithammer/shortuuid/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"grpc-dtm-workflow-test/rpc/orders"
	"grpc-dtm-workflow-test/rpc/pay"
	oders2 "grpc-dtm-workflow-test/services/orders"
	pay2 "grpc-dtm-workflow-test/services/pay"
	"time"

	"log"
	"net"
)

const PayAddr = "192.168.94.1:10001"
const OrderAddr = "192.168.94.1:10000"
const DTMAddr = "192.168.40.129:36790"

func main() {

	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(dtmgimp.GrpcServerLog))
	workflow.InitGrpc(DTMAddr, PayAddr, rpcServer)
	go func() {
		rpcServer.RegisterService(&pay.PayService_ServiceDesc, &pay2.Pay{})

		listen, err := net.Listen("tcp", PayAddr)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		log.Println("pay 服务启动成功 ", PayAddr)
		rpcServer.Serve(listen)
	}()
	go func() {
		rpcServer.RegisterService(&orders.OrdersService_ServiceDesc, &oders2.Order{})

		listen, err := net.Listen("tcp", OrderAddr)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		log.Println("order 服务启动成功 ", OrderAddr)
		rpcServer.Serve(listen)
	}()
	conn1, err := grpc.Dial(PayAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(workflow.Interceptor),
	)
	if err != nil {
		panic(err)
	}
	conn2, err := grpc.Dial(OrderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(workflow.Interceptor))
	if err != nil {
		panic(err)
	}
	payClient := pay.NewPayServiceClient(conn1)
	orderClient := orders.NewOrdersServiceClient(conn2)

	wfName := "test-grpc"
	err = workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
		// 创建订单
		var req2 orders.OrdersRequest
		req2.Id = 1
		req2.Amount = 120
		req2.UserId = 1
		wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err := orderClient.CreateRevert(wf.Context, &req2)
			log.Println("order revert", err)
			return err
		})
		_, err = orderClient.Create(wf.Context, &req2)
		if err != nil {
			log.Println("order", err)
			return err
		}
		// 支付
		var req pay.PayRequest
		err := proto.Unmarshal(data, &req)
		if err != nil {
			return err
		}
		wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err := payClient.PayRevert(wf.Context, &req)
			log.Println("pay revert", err)
			return err
		})
		_, err = payClient.Pay(wf.Context, &req)
		if err != nil {
			log.Println("pay", err)
			return err
		}
		return err
	})
	payReq := &pay.PayRequest{
		UserId: 1,
		Price:  120,
	}
	data, err := proto.Marshal(payReq)
	if err != nil {
		panic(err)
	}
	_, err = workflow.ExecuteCtx(context.Background(), wfName, shortuuid.New(), data)
	if err != nil {
		if errors.Is(err, dtmcli.ErrFailure) {
			log.Println("rollback")
			return
		}
		panic(err)
	}
	time.Sleep(3 * time.Second)
}
