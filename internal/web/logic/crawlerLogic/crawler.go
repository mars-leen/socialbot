package crawlerLogic

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/wblogger"
)

func Add(form *model.CrawlerForm) common.Result {
	crawler := model.Crawler{
		Name: form.Name,
		Script:form.Script,
		LastPage:form.LastPage,
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

	_,err = crawler.UpdateColsById(form.Id, "script", "name", "last_page")
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

func ListItem(crwid int) common.Result {
	var err error
	list := model.CrawlerItemList{}
	if crwid == 0 {
		err = list.GetList()
	}else {
		err = list.GetListByCrwid(crwid)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}
