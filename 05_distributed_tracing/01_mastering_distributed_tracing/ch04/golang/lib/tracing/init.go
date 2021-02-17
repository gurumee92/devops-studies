package tracing

import (
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Init is
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(
		service,
		config.Logger(jaeger.StdLogger),
	)

	if err != nil {
		log.Fatalf("Error: cannot init Jaeger: %v\n", err)
	}

	return tracer, closer
}
