package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		header := c.Request.Header
		headerBts, _ := json.Marshal(header)
		c.String(http.StatusOK, "%s", headerBts)
	})
	g.Run(":1234")
}
