package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
	"strconv"
	"strings"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		applyCORS(c)
		c.Next()
	}
}

func applyCORS(c *gin.Context )  {
	cfg,err := configService.GetCors()
	if err != nil {
		wblogger.Log.Error(err)
		return
	}
	if cfg.Enable == false {
		return
	}
	method := c.Request.Method
	if len(cfg.AllowOrigins) > 0 {
		c.Header("Access-Control-Allow-Origin", strings.Join(cfg.AllowOrigins, ","))
	}
	if len(cfg.AllowHeaders) > 0 {
		c.Header("Access-Control-Allow-Headers", strings.Join(cfg.AllowHeaders, ","))
	}
	if len(cfg.AllowMethods) > 0 {
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	}
	if cfg.AllowCredentials{
		c.Header("Access-Control-Allow-Credentials", "true")
	}else{
		c.Header("Access-Control-Allow-Credentials", "false")
	}
	if cfg.MaxAge > 0 {
		value := strconv.Itoa(cfg.MaxAge)
		c.Header("Access-Control-Max-Age", value)
	}
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}