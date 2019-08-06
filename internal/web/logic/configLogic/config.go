package configLogic

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/wblogger"
)

func Add(form *model.ConfigForm) common.Result {
	config := model.Config{
		Key: form.Key,
		Title:form.Title,
		Value:form.Value,
	}
	_, err := config.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESS(nil)
}

func Delete(id int) common.Result {
	config := model.Config{}
	isExist, err := config.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	config.IsDel = 1
	_,err = config.UpdateColsById(id, "is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Update(form *model.ConfigForm) common.Result {
	config := model.Config{}
	isExist, err := config.GetOneById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	config.Key = form.Key
	config.Title = form.Title
	config.Value = form.Value
	_,err = config.UpdateById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func List(key string) common.Result {
	list := model.ConfigList{}
	var err error
	if key=="" {
		err = list.GetList()
	}else{
		err = list.GetListByKey(key)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}
