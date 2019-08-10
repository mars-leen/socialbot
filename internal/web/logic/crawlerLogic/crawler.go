package crawlerLogic

import (
	"fmt"
	"net/url"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
	"strings"
)

func Add(form *model.CrawlerForm) common.Result {
	crawler := model.Crawler{
		Name:     form.Name,
		Script:   form.Script,
		LastPage: form.LastPage,
	}

	_, err := crawler.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Update(form *model.CrawlerForm) common.Result {
	crawler := model.Crawler{}
	isExist, err := crawler.GetOneById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	crawler.Script = form.Script
	crawler.Name = form.Name
	crawler.LastPage = form.LastPage

	_, err = crawler.UpdateColsById(form.Id, "script", "name", "last_page")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func List() common.Result {
	var err error
	list := model.TagList{}
	err = list.GetList()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}

func ListRandItem(crwid int) common.Result{
	conList := model.CrawlerItemList{}
	err := conList.GetRandList(crwid, common.CrawlerItemNew, common.CrawlerListPageNum)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// format
	l := len(conList)
	conMediaList := make([]*model.ConCrawlerItem, 0, l)
	for _, con := range conList {
		conMediaList = append(conMediaList, &model.ConCrawlerItem{
			Id:con.Id,
			Medias:FormatContent(con.Content),
			Title:con.Title,
			Cover:con.Cover,
			Description:con.Description,
			Crwid:con.Crwid,
		})
	}
	return common.SUCCESS(conMediaList)
}


func DeleteItem(id int64) common.Result {
	CrawlerItem := model.CrawlerItem{}
	isExist, err := CrawlerItem.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	CrawlerItem.IsDel = 1
	_, err = CrawlerItem.UpdateColsById(id, nil,"is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESS(nil)
}

func FormatContent(content string) []string {
	list := strings.Split(content, ",")
	for i,v := range list  {
		u,_ := url.Parse(v)
		list[i] = fmt.Sprintf("%s/v1/adminApi/reverse?scheme=%s&host=%s&path=%s&param=%s",configService.GetHost(),u.Scheme,u.Host,u.Path, "imageMogr2/thumbnail/!88p")
	}
	return list
}
