package main

import (
	"flag"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var tag = ""
var agentHostIP = ""

func init() {
	flag.StringVar(&tag, "tag", "tag", "tag")
	flag.Parse()
}
func main() {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		panic(err)
	}
	options := make([]jaegercfg.Option, 0, 3)
	options = append(options, jaegercfg.Logger(jaeger.StdLogger))
	tracer, TraceClose, err := cfg.NewTracer(options...)
	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)
	for {
		span := opentracing.StartSpan("test")
		span.SetTag("tag", tag)
		span.SetTag("tag1", "tag1value")
		span.LogKV("log1", "log1value")
		childSpan := opentracing.StartSpan("test-child", opentracing.ChildOf(span.Context()))
		childSpan.SetTag("tag2", "tag2value")
		childSpan.LogKV("log2", "log2value")
		time.Sleep(time.Second)
		childSpan.Finish()
		span.Finish()
		time.Sleep(time.Second * 3)
	}
	TraceClose.Close()
}
