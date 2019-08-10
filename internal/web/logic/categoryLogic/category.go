package categoryLogic

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/uploadLogic"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/categoryService"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/storageService"
	"socialbot/internal/web/wblogger"
)


func Add(c *gin.Context, form *model.CategoryForm) common.Result {
	filename,err := uploadLogic.UploadCategoryCover(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	category := model.Category{
		Title:form.Title,
		ShortName:form.ShortName,
		Description:form.Description,
		Sort:form.Sort,
	}

	if filename != "" {
		category.Cover = filename
	}

	_, err = category.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Delete(id int) common.Result {
	category := model.Category{}
	isExist, err := category.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	category.IsDel = 1
	_,err = category.UpdateColsById(id, "is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Update(c *gin.Context, form *model.CategoryForm) common.Result {
	filename,err := uploadLogic.UploadCategoryCover(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}


	category := model.Category{}
	isExist, err := category.GetOneById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	category.Sort = form.Sort
	category.Description = form.Description
	category.Title = form.Title
	category.ShortName = form.ShortName
	if filename != "" {
		if category.Cover != "" {
			storageService.SyncRemoveSigFile("local", filepath.Join(configService.GetStorageUploadPath(), category.Cover))
		}
		category.Cover = filename
	}

	_,err = category.UpdateById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func List() common.Result {
	list := model.CategoryList{}
	err := list.GetList()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	for i,v := range list  {
		list[i].Cover = configService.GetUploadFullUrl(v.Cover)
	}
	return common.SUCCESSARR(list)
}

func ListWithTags() common.Result  {
	result,err := categoryService.GetCategoryWithTag()
	if err != nil {
		return common.SystemError
	}
	return common.SUCCESSARR(result)
}