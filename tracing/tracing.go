package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var gtracer opentracing.Tracer

func GetTracer() opentracing.Tracer {
	return gtracer
}

func InitJaeger() io.Closer {
	cfg := &config.Configuration{
		ServiceName: "Bush",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "127.0.0.1:6831",
			LogSpans:           true,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}
	gtracer = tracer
	return closer
}
