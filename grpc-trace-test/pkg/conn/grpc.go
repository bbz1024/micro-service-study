package conn

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: grpc
 * @Date:
 * @Desc: ...
 *
 */

func Conn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				ctx = opentracing.ContextWithSpan(ctx, opentracing.SpanFromContext(ctx))
				return invoker(ctx, method, req, reply, cc, opts...)
			},
			//注入trace
			grpc_opentracing.UnaryClientInterceptor(),
		),
	)
}
