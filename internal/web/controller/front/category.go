package front

import (
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/logic/categoryLogic"
)


func ListCategory(c *gin.Context) {
	categoryLogic.List().Out(c)
}

func ListCategoryWithTags(c *gin.Context) {
	categoryLogic.ListWithTags().Out(c)
}
