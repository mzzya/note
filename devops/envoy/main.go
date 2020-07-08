package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var port *int = flag.Int("port", 1234, "help message for flagname")

func main() {
	flag.Parse()
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		header := c.Request.Header
		header.Set("a-hostName", os.Getenv("HOSTNAME"))
		headerBts, _ := json.Marshal(header)
		c.String(http.StatusOK, "%s", headerBts)
	})
	g.Run(fmt.Sprintf(":%d", *port))
}
