package tracing

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
)

func init() {
	os.Setenv("JAEGER_AGENT_HOST", "192.168.40.129")
	os.Setenv("JAEGER_AGENT_PORT", "6831")
}

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	cfg.ServiceName = service                  // 服务名
	cfg.Sampler.Type = jaeger.SamplerTypeConst // 采样类型,const/remote/probabilistic/rateLimiting
	cfg.Sampler.Param = 1                      // 采样率
	cfg.Reporter.LogSpans = true               // 是否打印日志

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

// StartSpan 开启一个span，如果传入的span不为空，那么就是子span
func StartSpan(ctx context.Context, name string) opentracing.Span {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		span = opentracing.StartSpan(name)
	} else {
		spContext := span.Context()
		span = opentracing.StartSpan(name, opentracing.ChildOf(spContext))
	}
	return span
}

func RecordWithIP(span opentracing.Span, ip string) {
	span.SetTag("ip", ip)
	span.LogFields(
		log.String("event", "ip"),
		log.String("ip", ip),
	)

}

// RecordError 记录错误
func RecordError(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
}
