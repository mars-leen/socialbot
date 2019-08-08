package admin

import (
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/galleryLogic"
	"strconv"
)

func AddGalleryTag(c *gin.Context) {
	idStr := c.DefaultPostForm("msid", "0")
	cidStr := c.DefaultPostForm("cid", "0")
	tags := c.DefaultPostForm("tags", "")
	if tags == "" {
		common.ParamError.Out(c)
		return
	}
	id,_ := strconv.ParseInt(idStr, 10, 64)
	if id == 0  {
		common.ParamError.Out(c)
		return
	}
	cid,_ := strconv.Atoi(cidStr)
	galleryLogic.AddTags(id, cid, tags).Out(c)
}

func ListGallery(c *gin.Context) {
	sourceTypeStr := c.DefaultQuery("sourceType", "0")
	lastIdStr := c.DefaultQuery("lastId", "0")
	sortStr := c.DefaultQuery("sort", "0")

	lastId,_ := strconv.ParseInt(lastIdStr, 10, 64)
	sourceType,_ := strconv.Atoi(sourceTypeStr)
	sort,_ := strconv.Atoi(sortStr)
	galleryLogic.List(lastId, sort, sourceType).Out(c)
}
