package main

import (
	"fmt"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"grpc-prometheus-test/rpc/user"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"log"
	"net"
)

const (
	port        = ":50051"
	metricsAddr = "192.168.94.1:5001"
)

func main() {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
		grpcprom.WithServerCounterOptions(
			grpcprom.WithConstLabels(map[string]string{
				"service": "user",
				"version": "v1",
				"env":     "dev",
			}), grpcprom.WithSubsystem("user"),
		),
	)
	client := prometheus.NewPedanticRegistry()
	client.MustRegister(srvMetrics)

	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(srvMetrics.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(srvMetrics.StreamServerInterceptor()),
	)
	//注册服务
	user.RegisterUserServiceServer(rpcServer, &User{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	httpSrv := &http.Server{Addr: metricsAddr}
	go func() {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.HandlerFor(
			client,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
		httpSrv.Handler = m
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("user 服务启动成功 http://127.0.0.1:50051")
	rpcServer.Serve(listen)
}
