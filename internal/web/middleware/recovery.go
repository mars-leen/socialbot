package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/wblogger"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				wblogger.Log.Error(errors.Wrap(err.(error), "request recovery"))
				c.JSON(http.StatusOK, common.SystemError)
			}
		}()
		c.Next()
	}
}