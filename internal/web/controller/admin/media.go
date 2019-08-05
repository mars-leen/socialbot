package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/mediaLogic"
	"socialbot/internal/web/model"
)

func AddCommissionProduct(c *gin.Context) {
	var form model.CommissionProductForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("%+v \n ,err %v", form, err)
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := mediaLogic.AddCommissionProduct(&form)
	c.JSON(http.StatusOK, result)
}

func AddSocialMediaFromCrawler(c *gin.Context) {
	var form model.SocialProductForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := mediaLogic.AddSocialMediaFromCrawler(&form)
	c.JSON(http.StatusOK, result)
}