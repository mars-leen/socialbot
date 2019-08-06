package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/robotServerLogic"
	"strconv"
)

func AddRobotServer(c *gin.Context) {
	idStr := c.DefaultPostForm("cid", "0")
	name := c.DefaultPostForm("name", "")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	if name == ""  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := robotServerLogic.Add(id, name)
	result.Out(c)
}

func DeleteRobotServer(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := robotServerLogic.Delete(id)
	c.JSON(http.StatusOK, result)
}

func ListRobotServer(c *gin.Context) {
	c.JSON(http.StatusOK, robotServerLogic.List())
}
