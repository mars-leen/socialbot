package configLogic

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
)

func Add(form *model.ConfigForm) common.Result {
	config := model.Config{
		KeyMark: form.Key,
		Title:   form.Title,
		Value:   form.Value,
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
	_, err = config.UpdateColsById(id, "is_del")
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

	config.KeyMark = form.Key
	config.Title = form.Title
	config.Value = form.Value
	_, err = config.UpdateById(form.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func UpdateByKey(form *model.ConfigForm) common.Result {
	var err error
	switch form.Key {
	case configService.WebsiteKey:
		_, err = configService.UpdateWebsite(form.Value)
		break
	case configService.CorsKey:
		_, err = configService.UpdateCors(form.Value)
		break
	case configService.ServerKey:
		_, err = configService.UpdateStorage(form.Value)
		break
	default:
		return common.ParamError
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESS(nil)
}

func List(key string) common.Result {
	list := model.ConfigList{}
	var err error
	if key == "" {
		err = list.GetList()
	} else {
		err = list.GetListByKey(key)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}

func Base() common.Result {
	website, err := configService.GetWebsite()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	cors, err := configService.GetCors()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	storage, err := configService.GetStorage()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(map[string]interface{}{
		"website": website,
		"cors": cors,
		"storage": storage,
	})
}
