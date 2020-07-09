package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hellojqk/simple/pkg/util"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

var port *int = flag.Int("port", 1234, "help message for flagname")
var serviceName = "envoy_app"
var operationName = "app"

func init() {
	//日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//链路跟踪配置
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		panic(err)
	}
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = "jaeger:6831"
	options := make([]jaegercfg.Option, 0, 4)
	options = append(options, jaegercfg.Logger(jaeger.StdLogger))

	//zipkin支持
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	options = append(options, jaegercfg.Injector(opentracing.HTTPHeaders, zipkinPropagator))
	options = append(options, jaegercfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator))
	options = append(options, jaegercfg.ZipkinSharedRPCSpan(true))

	tracer, _, err := cfg.NewTracer(options...)
	if err != nil {
		panic(err)
	}
	util.PrintJSONWithColor(cfg)
	opentracing.SetGlobalTracer(tracer)
}

func main() {
	flag.Parse()
	// gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		//测试日志
		uid, _ := uuid.NewRandom()
		uidStr := uid.String()
		log.Debug().
			Str("req_id", uidStr).
			Str("X-Request-Id", c.GetHeader("X-Request-Id")).
			Str("TraceId", c.GetHeader("X-B3-Traceid")).
			Str("ServerName", serviceName).
			Str("host", os.Getenv("HOSTNAME")).Msg("debug msg")
		//链路跟踪
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		var span opentracing.Span
		if spanCtx == nil {
			span = opentracing.StartSpan(operationName)
		} else {
			span = opentracing.StartSpan(operationName, opentracing.ChildOf(spanCtx))
		}
		defer span.Finish()
		span.SetTag("tag1", "tag1value")
		span.SetTag("app_req_id", uidStr)
		span.SetTag("app_X-Request-Id", c.GetHeader("X-Request-Id"))
		span.LogKV("log1", "log1value")
		//输出
		if err != nil {
			c.String(http.StatusBadRequest, "error:%s", err)
			return
		}
		header := c.Request.Header
		header.Set("a-hostName", os.Getenv("HOSTNAME"))
		headerBts, _ := json.Marshal(header)
		c.String(http.StatusOK, "%s", headerBts)
	})
	g.Run(fmt.Sprintf(":%d", *port))
}
