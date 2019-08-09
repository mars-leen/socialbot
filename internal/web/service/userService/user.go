package userService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
)

func MustGetTokenUser(c *gin.Context) (*model.User,error) {
	value, exist := c.Get(common.JwtAuthUserKey)
	if !exist {
		return nil, errors.New(fmt.Sprintf("GetTokenUser not exist, key %s", common.JwtAuthUserKey))
	}

	user, ok := value.(model.User)
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetTokenUser value.(models.User) value %v \n", value))
	}
	return &user, nil
}

func GetTokenUserId(c *gin.Context) (int64,error) {
	value, exist := c.Get(common.JwtAuthUserKey)
	if !exist {
		return 0, nil
	}

	user, ok := value.(model.User)
	if !ok {
		return 0, errors.New(fmt.Sprintf("GetTokenUser value.(models.User) value %v \n", value))
	}
	return user.Id, nil
}