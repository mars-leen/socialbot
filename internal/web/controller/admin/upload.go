package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/logic/uploadLogic"
)

func UploadSingle(c *gin.Context) {
	c.JSON(http.StatusOK,uploadLogic.UploadSingle(c, "file"))
}
