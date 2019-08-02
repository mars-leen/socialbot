package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/userLogic"
	"socialbot/internal/web/model"
)

func Login(c *gin.Context) {
	var accountForm model.AccountForm
	if err := c.ShouldBind(&accountForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := userLogic.AdminLogin(accountForm)
	c.JSON(http.StatusOK, result)
}

