package tagLogic

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/categoryService"
	"socialbot/internal/web/wblogger"
)

func Add(form *model.TagForm) common.Result {
	tag := model.Tag{
		Cid:         form.Cid,
		Title:       form.Title,
		ShortName:   form.ShortName,
		Description: form.Description,
		BoardName:   form.BoardName,
	}

	_, err := tag.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Delete(id int) common.Result {
	tag := model.Tag{}
	isExist, err := tag.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	tag.IsDel = 1
	_, err = tag.UpdateColsById(id, "is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Update(form *model.TagForm) common.Result {
	tag := model.Tag{}
	isExist, err := tag.GetOneById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	tag.BoardName = form.BoardName
	tag.Description = form.Description
	tag.Title = form.Title
	tag.ShortName = form.ShortName
	tag.Cid = form.Cid

	_, err = tag.UpdateById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func List(cid int) common.Result {
	var err error
	list := model.TagList{}
	if cid == 0 {
		err = list.GetList()
	} else {
		err = list.GetListByCid(cid)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	categoryMap, err := categoryService.GetCategoryMapName()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	tagList := make([]model.ConTag, len(list))
	for i, v := range list {
		categoryName, ok := categoryMap[v.Cid]
		if !ok {
			categoryName = ""
		}
		tagList[i] = model.ConTag{
			Id:           v.Id,
			Cid:          v.Cid,
			Title:        v.Title,
			BoardName:    v.BoardName,
			Description:  v.Description,
			ShortName:    v.ShortName,
			CategoryName: categoryName,
		}
	}
	return common.SUCCESSARR(list)
}
