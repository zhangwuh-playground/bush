package tracing

import (
	"fmt"
	"io"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var gtracer opentracing.Tracer

func GetTracer() opentracing.Tracer {
	return gtracer
}

func InitJaeger() io.Closer {
	host := os.Getenv("JAEAGER_COLLECTOR_ADDR")
	if len(host) == 0 {
		host = "127.0.0.1:6831"
	}
	cfg := &config.Configuration{
		ServiceName: "Bush",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: host,
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}
	gtracer = tracer
	return closer
}
