package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/mediaLogic"
	"strconv"
)

func MediaDetail(c *gin.Context)  {
	uriString := c.DefaultPostForm("uri", "0")
	uri, _ := strconv.ParseInt(uriString, 10, 64)
	if uri == 0 {
		common.ParamError.Out(c)
		return
	}
	mediaLogic.Detail(c, uri).Out(c)
}

func ListMedias(c *gin.Context)  {
	mediaLogic.ListByRecommend(0, 3).Out(c)
}

func HomeRecommendMedias(c *gin.Context)  {
	lastIdString := c.DefaultQuery("lastId", "0")
	CategoryString := c.DefaultQuery("category", "0")
	SortString := c.DefaultQuery("sort", "0")
	lastId, _ := strconv.ParseInt(lastIdString, 10, 64)
	category, _ := strconv.Atoi(CategoryString)
	sort, _ := strconv.Atoi(SortString)
	mediaLogic.ListByCategory(lastId, category, sort).Out(c)
}



func LikeMedia(c *gin.Context)  {
	uriString := c.DefaultPostForm("uri", "0")
	isLikeString := c.DefaultPostForm("isLike", "0")
	uri, _ := strconv.ParseInt(uriString, 10, 64)
	isLike, _ := strconv.Atoi(isLikeString)
	if uri == 0 {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	delStatus := 0
	if isLike != 1{
		delStatus = 1
	}
	mediaLogic.Like(c, uri,delStatus).Out(c)
}
