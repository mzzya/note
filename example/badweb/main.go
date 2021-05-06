package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var g = gin.New()

type RequestURI struct {
	Second     int `json:"second" uri:"second" form:"second"`
	StatusCode int `json:"status_code" uri:"status_code" form:"status_code"`
}

func main() {
	var port int
	flag.IntVar(&port, "p", 8899, "端口号，默认为8899")
	flag.Parse()

	g.GET("/s", func(c *gin.Context) {
		c.JSON(http.StatusOK, "哈哈哈")
	})

	timeOutGroup := g.Group("/timeout")

	timeOutGroup.GET("/:second", func(c *gin.Context) {
		var m = RequestURI{}
		c.BindUri(&m)
		time.Sleep(time.Duration(m.Second * int(time.Second)))
		c.JSON(http.StatusOK, m)
	})

	errorGroup := g.Group("/error")
	errorGroup.GET("/:status_code", func(c *gin.Context) {
		var m = RequestURI{}
		c.BindUri(&m)
		c.JSON(m.StatusCode, m)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: g,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("Server exiting")
}
