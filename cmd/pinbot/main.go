package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	g := gin.Default()
	g.GET("/api", func(c *gin.Context) {
		a := map[string]interface{}{
			"ip":c.ClientIP(),
			"hd":c.Request.Header,
		}
		c.JSON(http.StatusOK,a)
	})
	g.Run(":8877")
}
