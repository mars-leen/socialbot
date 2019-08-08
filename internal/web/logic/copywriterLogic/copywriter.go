package CopywriterLogic

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/wblogger"
)


func Add(form *model.CopywriterForm) common.Result {
	Copywriter := model.Copywriter{
		Title:form.Title,
		Description:form.Description,
	}
	_, err := Copywriter.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Delete(id int) common.Result {
	Copywriter := model.Copywriter{}
	isExist, err := Copywriter.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	Copywriter.IsDel = 1
	_,err = Copywriter.UpdateColsById(id, "is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Update(form *model.CopywriterForm) common.Result {
	Copywriter := model.Copywriter{}
	isExist, err := Copywriter.GetOneById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	Copywriter.Description = form.Description
	Copywriter.Title = form.Title
	_,err = Copywriter.UpdateById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func List() common.Result {
	list := model.CopywriterList{}
	err := list.GetList()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}
