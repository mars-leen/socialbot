package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/tagLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func AddTag(c *gin.Context) {
	var TagForm model.TagForm
	if err := c.ShouldBind(&TagForm); err != nil {
		fmt.Println(TagForm, err)
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := tagLogic.Add(&TagForm)
	c.JSON(http.StatusOK, result)
}

func UpdateTag(c *gin.Context) {
	var TagForm model.TagForm
	if err := c.ShouldBind(&TagForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := tagLogic.Update(&TagForm)
	c.JSON(http.StatusOK, result)
}

func DeleteTag(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := tagLogic.Delete(id)
	c.JSON(http.StatusOK, result)
}

func ListTag(c *gin.Context) {
	idStr := c.DefaultQuery("cid", "0")
	cid,_ := strconv.Atoi(idStr)
	c.JSON(http.StatusOK, tagLogic.List(cid))
}



