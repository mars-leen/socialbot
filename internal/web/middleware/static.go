package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"socialbot/internal/web/setting"
	"strings"
)

func ServeThemeView() gin.HandlerFunc {
	return func(c *gin.Context) {
		// filter api
		if strings.Contains(c.Request.URL.Path, "v1") {
			return
		}
		// admin theme
		if strings.Contains(c.Request.URL.Path, "dashboard") {
			dashboardPath := filepath.Join(setting.AppPath, "views/admin/socailbot-brain/dist")
			filePath := filepath.Join( dashboardPath, filepath.FromSlash(strings.Replace(c.Request.URL.Path, "/dashboard", "", 1)))
			_, err := os.Stat(filePath)
			if err != nil && os.IsNotExist(err) {
				filePath = filepath.Join(dashboardPath, "index.html")
			}
			http.ServeFile(c.Writer, c.Request, filePath)
			c.Abort()
			return
		}

		// front theme
		frontPath := filepath.Join(setting.AppPath, "views/front/socailbot-face/dist")
		filePath := filepath.Join( frontPath, filepath.FromSlash(c.Request.URL.Path))
		_, err := os.Stat(filePath)
		if err != nil && os.IsNotExist(err) {
			filePath = filepath.Join(frontPath, "index.html")
		}
		http.ServeFile(c.Writer, c.Request, filePath)
		c.Abort()
		return
	}
}

