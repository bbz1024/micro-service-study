package main

import (
	"context"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-limit-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

/*
限流：
 - 基于中间件限流器（令牌桶）
 基于grpc-middleware自定义限流器
 基于sentinel的限流

*/

func NewLimiterMiddleware(r *rate.Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r.Allow() {
			return handler(ctx, req)
		}
		return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
	}
}

type MyLimiter struct {
	r *rate.Limiter
}

func (l MyLimiter) Limit() bool {
	if l.r.Allow() {
		return true
	}
	return false
}
func main() {

	var interceptor []grpc.UnaryServerInterceptor
	/*
		//令牌桶限流
		interceptor = append(interceptor, NewLimiterMiddleware(
			rate.NewLimiter(rate.Limit(2), 5), // 每1秒生成2个令牌，最多存储10个令牌
		))
		//grpc-middleware自定义限流器
		var limiter MyLimiter
		limiter.r = rate.NewLimiter(rate.Limit(2), 5)
		interceptor = append(interceptor, ratelimit.UnaryServerInterceptor(limiter))

	*/
	//sentinel限流 配置一条限流规则
	//初始化sentinel
	if err := sentinel.InitDefault(); err != nil {
		panic(err)
	}
	_, err := flow.LoadRules([]*flow.Rule{
		{
			Resource:               "get_user",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              5,
		},
	})
	if err != nil {
		panic(err)
	}
	grpc.ChainUnaryInterceptor(interceptor...)
	rpcServer := grpc.NewServer()
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
