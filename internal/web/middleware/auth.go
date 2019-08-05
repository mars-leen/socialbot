package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/jwtauth"
	"strconv"
)

func UserInfo()  gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		claims, err := jwtauth.ParseToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		userId, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			wblogger.Log.Error(err)
			c.Next()
			return
		}

		if userId == 0 {
			c.Next()
			return
		}

		user := model.User{}
		exist, err := user.GetOneById(userId)
		if err != nil {
			wblogger.Log.Error(err)
			c.Next()
			return
		}
		if !exist {
			c.JSON(http.StatusOK, common.AuthError)
			c.Abort()
			return
		}

		c.Set(common.JwtAuthUserKey, user)
		c.Next()
	}
}

func Auth()  gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exist := c.Get(common.JwtAuthUserKey)
		if !exist {
			c.JSON(http.StatusOK, common.AuthError)
			c.Abort()
			return
		}

		_, ok := value.(model.User)
		if !ok {
			wblogger.Log.Error(errors.Errorf("Auth user value.(models.User) value %v \n", value))
			c.JSON(http.StatusOK, common.SystemError)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthAdmin()  gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		claims, err := jwtauth.ParseToken(tokenString)
		authError := common.AuthError
		if err != nil {
			authError.Out(c)
			c.Abort()
			return
		}

		userId, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			wblogger.Log.Error(errors.Wrap(err, fmt.Sprintf("strconv.ParseInt(%+v) failed\n", claims)))
			authError.Out(c)
			c.Abort()
			return
		}

		if userId == 0 {
			c.JSON(http.StatusOK, common.AuthError)
			c.Abort()
			return
		}

		user := model.User{}
		exist, err := user.GetOneById(userId)
		if err != nil {
			wblogger.Log.Error(err)
			authError.Out(c)
			c.Abort()
			return
		}
		if !exist {
			authError.Out(c)
			c.Abort()
			return
		}

		if user.Identity != common.SuperAdmin {
			authError.Out(c)
			c.Abort()
			return
		}

		c.Set(common.JwtAuthUserKey, user)
		c.Next()
	}
}
