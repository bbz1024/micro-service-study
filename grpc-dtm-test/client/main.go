package main

import (
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/lithammer/shortuuid/v3"
	"grpc-dtm-test/rpc/orders"
	"grpc-dtm-test/rpc/pay"
	"time"
)

const PayAddr = "192.168.94.1:10001"
const OrderAddr = "192.168.94.1:10000"
const DTMAddr = "192.168.40.129:36790"

func main() {
	time.Sleep(100 * time.Millisecond)
	orderID := 1
	orderReq := &orders.OrdersRequest{
		Id:     int32(orderID),
		Amount: 101,
		UserId: 1,
	}
	payReq := &pay.PayRequest{
		UserId: 1,
		Price:  101,
	}
	gid := shortuuid.New()
	saga := dtmgrpc.NewSagaGrpc(DTMAddr, gid).
		//这里会进行找到proto，其中需要指定package name
		Add(OrderAddr+"/orders.OrdersService/Create", OrderAddr+"/orders.OrdersService/CreateRevert", orderReq).
		Add(PayAddr+"/pay.PayService/Pay", PayAddr+"/pay.PayService/PayRevert", payReq)
	err := saga.Submit()
	if err != nil {
		panic(err)
	}
	//run()
	select {}
}

//func run() {
//	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(dtmgimp.GrpcServerLog))
//	workflow.InitGrpc(DTMAddr, Addr1, rpcServer)
//	go func() {
//		rpcServer.RegisterService(&pay.PayService_ServiceDesc, &pay2.Pay{})
//
//		listen, err := net.Listen("tcp", Addr1)
//		if err != nil {
//			fmt.Println(err)
//			panic(err)
//		}
//		log.Println("user 服务启动成功 ", Addr1)
//		rpcServer.Serve(listen)
//	}()
//
//	conn1, err := grpc.Dial(Addr1,
//		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(workflow.Interceptor),
//	)
//	if err != nil {
//		panic(err)
//	}
//	client := pay.NewPayServiceClient(conn1)
//	wfName := "test-grpc"
//	err = workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
//		var req pay.PayRequest
//		err := proto.Unmarshal(data, &req)
//		if err != nil {
//			log.Println(err)
//			panic(err)
//		}
//		wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
//			_, err := client.PayRevert(wf.Context, &req)
//			return err
//		})
//		_, err = client.Pay(wf.Context, &req)
//
//		return err
//	})
//	fmt.Println(err)
//	payReq := &pay.PayRequest{
//		UserId: 1,
//		Price:  88,
//	}
//	data, err := proto.Marshal(payReq)
//	if err != nil {
//		panic(err)
//	}
//	_, err = workflow.ExecuteCtx(context.Background(), wfName, shortuuid.New(), data)
//	if err != nil {
//		panic(err)
//	}
//	time.Sleep(3 * time.Second)
//}
