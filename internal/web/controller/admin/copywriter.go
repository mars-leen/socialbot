package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/CopywriterLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func AddCopywriter(c *gin.Context) {
	var CopywriterForm model.CopywriterForm
	if err := c.ShouldBind(&CopywriterForm); err != nil {
		fmt.Println(CopywriterForm)
		common.ParamError.Out(c)
		return
	}
	result := CopywriterLogic.Add(&CopywriterForm)
	result.Out(c)
}

func UpdateCopywriter(c *gin.Context) {
	var CopywriterForm model.CopywriterForm
	if err := c.ShouldBind(&CopywriterForm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := CopywriterLogic.Update(&CopywriterForm)
	c.JSON(http.StatusOK, result)
}

func DeleteCopywriter(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := CopywriterLogic.Delete(id)
	c.JSON(http.StatusOK, result)
}

func ListCopywriter(c *gin.Context) {
	c.JSON(http.StatusOK, CopywriterLogic.List())
}



