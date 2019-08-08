package admin

import (
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/mediaLogic"
	"socialbot/internal/web/model"
)

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

