package admin

import (
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/mediaLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func MediaDetail(c *gin.Context)  {
	uriString := c.DefaultQuery("uri", "0")
	uri, _ := strconv.ParseInt(uriString, 10, 64)
	if uri == 0 {
		common.ParamError.Out(c)
		return
	}
	mediaLogic.Detail(c, uri).Out(c)
}

func EditMedia(c *gin.Context)  {
	var form model.EditMediaForm
	if err := c.ShouldBind(&form); err != nil {
		common.ParamError.Errorf("%v", err).Out(c)
		return
	}
	mediaLogic.Edit(&form).Out(c)
}

func ListMedias(c *gin.Context)  {
	lastIdString := c.DefaultQuery("lastId", "0")
	CategoryString := c.DefaultQuery("category", "0")
	SortString := c.DefaultQuery("sort", "0")
	lastId, _ := strconv.ParseInt(lastIdString, 10, 64)
	category, _ := strconv.Atoi(CategoryString)
	sort, _ := strconv.Atoi(SortString)
	mediaLogic.ListByCategory(lastId, category, sort).Out(c)
}


func AddCommissionProduct(c *gin.Context) {
	var form model.CommissionProductForm
	if err := c.ShouldBind(&form); err != nil {
		common.ParamError.Errorf("%v", err).Out(c)
		return
	}
	mediaLogic.AddCommissionProduct(&form).Out(c)
}

func AddSocialMediaFromCrawler(c *gin.Context) {
	var form model.SocialProductForm
	if err := c.ShouldBind(&form); err != nil {
		common.ParamError.Errorf("%v", err).Out(c)
		return
	}
	mediaLogic.AddSocialMediaFromCrawler(&form).Out(c)
}

