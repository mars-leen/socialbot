package front

import (

	"github.com/gin-gonic/gin"
	"net/http"

	"socialbot/internal/web/logic/tagLogic"

	"strconv"
)

func ListTag(c *gin.Context) {
	idStr := c.DefaultQuery("cid", "0")
	cid,_ := strconv.Atoi(idStr)
	c.JSON(http.StatusOK, tagLogic.List(cid))
}



