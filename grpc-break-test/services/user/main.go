package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
	"google.golang.org/grpc"
	"grpc-breake-test/rpc/user"
	"log"

	"net"
)

const (
	port = ":50051"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {

	fmt.Println("熔断关闭")
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {

	fmt.Println("触发熔断")
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Println("熔断恢复")
}

var (
	slowRules = []*circuitbreaker.Rule{
		{
			Resource:                     "get_user",
			Strategy:                     circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:               3000, // 重试超时时间
			MinRequestAmount:             5,    // 最小请求数
			StatIntervalMs:               5000, // 统计时间间隔
			StatSlidingWindowBucketCount: 10,   // 滑动窗口数
			MaxAllowedRtMs:               40,   // 最大响应时间
			Threshold:                    0.5,  // 50% 错误率
		},
	}
	errorRateRules = []*circuitbreaker.Rule{
		{
			Resource:                     "get_user",
			Strategy:                     circuitbreaker.ErrorRatio,
			RetryTimeoutMs:               3000,
			MinRequestAmount:             10,
			StatIntervalMs:               5000,
			StatSlidingWindowBucketCount: 10,
			Threshold:                    0.4, // 40% 错误率
		},
	}
	// Statistic time span=5s, recoveryTimeout=3s, maxErrorCount=5
	errorCountRules = []*circuitbreaker.Rule{
		{
			Resource:                     "get_user",
			Strategy:                     circuitbreaker.ErrorCount,
			RetryTimeoutMs:               3000,
			MinRequestAmount:             10,
			StatIntervalMs:               5000,
			StatSlidingWindowBucketCount: 10,
			Threshold:                    5, // 5次错误 10次请求其中有5次错误，触发熔断
		},
	}
)

func main() {
	//熔断
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	//注册熔断监听
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})
	//基于慢调用比例的熔断
	//_, err = circuitbreaker.LoadRules(slowRules)
	//基于错误率的熔断
	//_, err = circuitbreaker.LoadRules(errorRateRules)
	//基于异常数的熔断
	_, err = circuitbreaker.LoadRules(errorCountRules)
	if err != nil {
		log.Fatal(err)
	}
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
