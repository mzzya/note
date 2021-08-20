package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var port *int = flag.Int("port", 10086, "help message for port")

func main() {
	flag.Parse()
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		header := c.Request.Header
		header.Set("a-hostName", os.Getenv("HOSTNAME"))
		// headerBts, _ := json.Marshal(header)
		c.JSON(http.StatusOK, gin.H{"header": header,
			//Request不能直接序列化
			// "Request":    c.Request,
			"URL":        c.Request.URL,
			"RemoteAddr": c.Request.RemoteAddr,
			"RequestURI": c.Request.RequestURI,
			"ClientIP":   c.ClientIP(),
		})
	})
	g.Run(fmt.Sprintf(":%d", *port))
}
