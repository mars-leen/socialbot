package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func FrontView(root string) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		fmt.Println("---->",p)
		fpath := filepath.Join(root, filepath.FromSlash(p))
		_, verr := os.Stat(fpath)
		if verr != nil && os.IsNotExist(verr) {
			fpath = filepath.Join(root, "index.html")
		}
		http.ServeFile(c.Writer, c.Request, fpath)
		c.Abort()
	}
}
