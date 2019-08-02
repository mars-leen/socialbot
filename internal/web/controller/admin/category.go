package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/categoryLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func AddCategory(c *gin.Context) {
	var categoryForm model.CategoryForm
	if err := c.ShouldBind(&categoryForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := categoryLogic.Add(&categoryForm)
	c.JSON(http.StatusOK, result)
}

func UpdateCategory(c *gin.Context) {
	var categoryForm model.CategoryForm
	if err := c.ShouldBind(&categoryForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := categoryLogic.Update(&categoryForm)
	c.JSON(http.StatusOK, result)
}

func DeleteCategory(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	if id == 0  {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := categoryLogic.Delete(id)
	c.JSON(http.StatusOK, result)
}

func ListCategory(c *gin.Context) {
	c.JSON(http.StatusOK, categoryLogic.List())
}

func ListCategoryWithTags(c *gin.Context) {
	c.JSON(http.StatusOK, categoryLogic.ListWithTags())
}
