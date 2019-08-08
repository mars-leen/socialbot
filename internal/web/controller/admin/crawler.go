package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/crawlerLogic"
	"socialbot/internal/web/model"
	"strconv"
)

func AddCrawler(c *gin.Context) {
	var CrawlerForm model.CrawlerForm
	if err := c.ShouldBind(&CrawlerForm); err != nil {
		common.ParamError.Errorf("%v", err).Out(c)
		return
	}
	crawlerLogic.Add(&CrawlerForm).Out(c)
}

func UpdateCrawler(c *gin.Context) {
	var CrawlerForm model.CrawlerForm
	if err := c.ShouldBind(&CrawlerForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	crawlerLogic.Update(&CrawlerForm).Out(c)
}


func ListCrawler(c *gin.Context) {
	crawlerLogic.List().Out(c)
}

func ListRandCrawlerItem(c *gin.Context) {
	idStr := c.DefaultPostForm("crwid", "0")
	crwid,_ := strconv.Atoi(idStr)
	crawlerLogic.ListRandItem(crwid).Out(c)
}

func DeleteCrawlerItem(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	crawlerLogic.DeleteItem(id).Out(c)
}