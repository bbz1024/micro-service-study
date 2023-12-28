package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"micro-service/service"
	"net"
	"net/http"
)

/*
	func main() {
		//添加证书
		creds, err := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
		if err != nil {
			fmt.Println(err)
		}

		rpcServer := grpc.NewServer(grpc.Creds(creds))
		service.RegisterProdServiceServer(rpcServer, service.ProductService)

		listen, err := net.Listen("tcp", ":8002")
		if err != nil {
			fmt.Println(err)
		}
		rpcServer.Serve(listen)
	}
*/
/*
func main() {
	//	双向认证
	cret, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cretPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    cretPool,
	})
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer.Serve(listen)
}


*/
/*
func main() {
	//	Token认证
	cret, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cretPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    cretPool,
	})
	Auth := func(ctx context.Context) error {
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			return fmt.Errorf("no metadata")
		}
		var user, password string
		if val, ok := md["user"]; ok {
			user = val[0]
		}
		if val, ok := md["password"]; ok {
			password = val[0]
		}
		if user != "admin" || password != "123456" {
			fmt.Println(222)
			return status.Errorf(codes.Unauthenticated, "token 不合法")
		}
		return nil
	}
	//基于拦截器jxToken认证
	rpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			err = Auth(ctx)
			fmt.Println(ctx)
			if err != nil {
				//进行拦截
				return
			}
			//正常执行
			return handler(ctx, req)
		}))
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer.Serve(listen)
}


*/
/*
//服务端流
func main() {
	//	Token认证
	cret, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cretPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    cretPool,
	})
	Auth := func(ctx context.Context) error {
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			return fmt.Errorf("no metadata")
		}
		var user, password string
		if val, ok := md["user"]; ok {
			user = val[0]
		}
		if val, ok := md["password"]; ok {
			password = val[0]
		}
		if user != "admin" || password != "123456" {
			fmt.Println(222)
			return status.Errorf(codes.Unauthenticated, "token 不合法")
		}
		return nil
	}
	//基于拦截器jxToken认证
	rpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			err = Auth(ctx)
			fmt.Println(ctx)
			if err != nil {
				//进行拦截
				return
			}
			//正常执行
			return handler(ctx, req)
		}))
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer.Serve(listen)
}


*/
/*
//双向流
func main() {
	cret, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cretPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    cretPool,
	})

	rpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor),
	)
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer.Serve(listen)
}
func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		return fmt.Errorf("no metadata")
	}
	var user, password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "123456" {
		fmt.Println(222)
		return status.Errorf(codes.Unauthenticated, "token 不合法")
	}
	return nil
}
func interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	err = Auth(ctx)
	fmt.Println("服务端拦截器")
	if err != nil {
		//进行拦截
		return
	}
	//正常执行
	return handler(ctx, req)
}


*/
// 内置trace

func main() {
	cret, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cretPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cret},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    cretPool,
	})

	rpcServer := grpc.NewServer(
		grpc.Creds(creds),
	)
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
	}
	go startTrace()
	rpcServer.Serve(listen)
}
func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	grpclog.Println("Trace listen on 50051")
}
