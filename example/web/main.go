package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var startTime = time.Now().Format("2006-01-02 15:04:05")
var ip = GetIP()

func main() {
	go WebHealth()
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "StartTime:%s\tIP:%s", startTime, ip)
	})
	srv := http.Server{Addr: ":8888", Handler: g}
	quit := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server listern err", err)
			quit <- os.Kill
		}
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("server shuddown error", err)
	}
	fmt.Println("server close")
}

// WebHealth web监听
func WebHealth() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Error(errors.WithMessage(err, "web_health"))
	}
}

// GetIP .
func GetIP() (ip string) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range addr {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return
			}
		}
	}
	return
}
