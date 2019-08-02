package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/wblogger"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		start := time.Now()
		c.Next()
		timeConsuming := time.Since(start).Nanoseconds() / 1e6
		wblogger.Request.Info( fmt.Sprintf(
			"path:%s method:%s ip:%s url:%s proto:%s res_status:%d res_length:%d timeConsuming:%d, user_agent:%s header:%+v",
			path,method,c.Request.Proto,c.Request.URL.String(),c.Request.Proto, c.Writer.Status(),c.Writer.Size(),timeConsuming, c.GetHeader("User-Agent"), c.Request.Header))
	}
}
