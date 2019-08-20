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

func BaseConfig(c *gin.Context) {
	configLogic.Base().Out(c)
}

func AddConfig(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	result := configLogic.Add(&form)
	result.Out(c)
}

func UpdateConfig(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		common.ParamError.Out(c)
		return
	}
	result := configLogic.UpdateByKey(&form)
	result.Out(c)
}

func ListConfig(c *gin.Context) {
	configLogic.List("").Out(c)
}

func AddServer(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	form.Key = configService.ServerKey
	configLogic.Add(&form).Out(c)
}

func DeleteServer(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		common.ParamError.Out(c)
		return
	}
	configLogic.Delete(id).Out(c)
}

func UpdateServer(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	configLogic.Update(&form).Out(c)

}

func ListServer(c *gin.Context) {
	configLogic.List(configService.ServerKey).Out(c)
}

func AddReverseHost(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	form.Key = configService.ReserveHostKey
	configLogic.Add(&form).Out(c)
}

func DeleteReverseHost(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		common.ParamError.Out(c)
		return
	}
	configLogic.Delete(id).Out(c)
}

func UpdateReverseHost(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	configLogic.Update(&form).Out(c)
}

func ListReverseHost(c *gin.Context) {
	configLogic.List(configService.ReserveHostKey).Out(c)
}
