package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/configLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func BaseConfig(c *gin.Context) {
	configLogic.Base().Out(c)
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

func AddServer(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	form.Key = common.ServerConfigKey
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
	configLogic.List(common.ServerConfigKey).Out(c)
}

func AddReverseHost(c *gin.Context) {
	var form model.ConfigForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(form, err)
		common.ParamError.Out(c)
		return
	}
	form.Key = common.ReserveHostConfigKey
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
	configLogic.List(common.ReserveHostConfigKey).Out(c)
}
