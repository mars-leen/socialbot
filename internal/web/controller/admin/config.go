package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/configLogic"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"strconv"
)

func AddServer(c *gin.Context)  {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	form.Key = configService.ServerKey
	result := configLogic.Add(&form)
	result.Out(c)
}

func DeleteServer(c *gin.Context)  {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		common.ParamError.Out(c)
		return
	}
	result := configLogic.Delete(id)
	result.Out(c)
}

func UpdateServer(c *gin.Context)  {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	result := configLogic.Update(&form)
	result.Out(c)
}

func ListServer(c *gin.Context) {
	configLogic.List(configService.ServerKey).Out(c)
}

