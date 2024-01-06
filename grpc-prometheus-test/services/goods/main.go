package main

import (
	"fmt"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"grpc-prometheus-test/rpc/goods"
	"log"
	"net/http"

	"net"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: main
 * @Date:
 * @Desc: ...
 *
 */
const (
	port        = ":50052"
	metricsAddr = "192.168.94.1:5001"
)

func main() {

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	client := prometheus.NewRegistry()
	client.MustRegister(srvMetrics)

	rpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(srvMetrics.UnaryServerInterceptor()),
		//grpc.ChainStreamInterceptor(srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(prom.ExtractContext))),
	)
	//注册服务
	goods.RegisterGoodsServiceServer(rpcServer, &Goods{})
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
		//if err := httpSrv.ListenAndServe(); err != nil {
		//	log.Println(err)
		//}
	}()
	log.Println("goods 服务启动成功 http://127.0.0.1:50052")
	rpcServer.Serve(listen)
}
