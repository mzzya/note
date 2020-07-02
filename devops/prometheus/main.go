package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var orderCounter *prometheus.CounterVec

func init() {
	var ops = prometheus.CounterOpts{
		Namespace: "business",
		Subsystem: "order",
		Name:      "total",
		Help:      "订单系统统计",
	}
	orderCounter = prometheus.NewCounterVec(ops, []string{"site", "type"})
	orderCounter.WithLabelValues("smb", "all").Add(100000)
	orderCounter.WithLabelValues("smb", "success").Add(99999)
	orderCounter.WithLabelValues("smb", "cancel").Add(1111)

}

func main() {

	go func() {
		for {
			orderCounter.WithLabelValues("smb", "all").Add(10)
			orderCounter.WithLabelValues("smb", "success").Add(9)
			orderCounter.WithLabelValues("smb", "cancel").Add(1)
			time.Sleep(time.Second)
		}
	}()

	prometheus.MustRegister(orderCounter)
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(":6060", nil))
}
