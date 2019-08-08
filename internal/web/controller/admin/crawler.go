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
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := crawlerLogic.Add(&CrawlerForm)
	c.JSON(http.StatusOK, result)
}

func UpdateCrawler(c *gin.Context) {
	var CrawlerForm model.CrawlerForm
	if err := c.ShouldBind(&CrawlerForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := crawlerLogic.Update(&CrawlerForm)
	c.JSON(http.StatusOK, result)
}


func ListCrawler(c *gin.Context) {
	c.JSON(http.StatusOK, crawlerLogic.List())
}

func ListRandCrawlerItem(c *gin.Context) {
	idStr := c.DefaultPostForm("crwid", "0")
	crwid,_ := strconv.Atoi(idStr)
	c.JSON(http.StatusOK, crawlerLogic.ListRandItem(crwid))
}

func DeleteCrawlerItem(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id,_ := strconv.Atoi(idStr)
	c.JSON(http.StatusOK, crawlerLogic.DeleteItem(id))
}