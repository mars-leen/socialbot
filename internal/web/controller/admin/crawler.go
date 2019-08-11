package admin

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/crawlerLogic"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
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
	id,_ := strconv.ParseInt(idStr, 10, 64)
	crawlerLogic.DeleteItem(id).Out(c)
}
//http://localhost:8080/v1/adminApi/reverse?scheme=https&host=qnm.hunliji.com&path=/o_1d8ke0df31ftk1tj81p49kh315r9p.jpg&param=imageMogr2/thumbnail/!88p
//http://localhost:8080/v1/adminApi/reverse?scheme=http&host=139.199.132.157:8877&path=/api
func Reverse(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		wblogger.Log.Error(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// you can reassign the body if you need to parse it as multipart
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a new url from the raw RequestURI sent by the client
	fullUrl := fmt.Sprintf("%s://%s%s?%s",c.Query("scheme"), c.Query("host"), c.Query("path"), c.Query("param"))
	if c.Query("param") != "" {
		fullUrl = fmt.Sprintf("%s://%s%s",c.Query("scheme"), c.Query("host"), c.Query("path"))
	}

	proxyReq, err := http.NewRequest(c.Request.Method, fullUrl, bytes.NewReader(body))
	if err != nil {
		wblogger.Log.Error(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	header := configService.GetHeader(c.Query("host"))
	for key, value := range header {
		proxyReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		wblogger.Log.Error(err)
		http.Error(c.Writer, err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()
	bodyContent, _ := ioutil.ReadAll(resp.Body)
	_, err = c.Writer.Write(bodyContent)
	if err != nil {
		wblogger.Log.Error(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	
	for h, val := range resp.Header {
		c.Writer.Header()[h] = val
	}
	c.Abort()
}