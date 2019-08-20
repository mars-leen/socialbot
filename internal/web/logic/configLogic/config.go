package configLogic

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
)

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
		"cors":    cors,
		"storage": storage,
	})
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

func Add(form *model.ConfigForm) common.Result {
	config := model.Config{
		KeyMark: form.Key,
		Title:   form.Title,
		Value:   form.Value,
	}

	// del cache 为了防止缓存穿透，首次访问数据库内不存在的host,会设置空值，此处需要删掉
	DelCache(config.KeyMark, config.Title)

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
	// del cache
	DelCache(config.KeyMark, config.Title)

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

	// del cache
	DelCache(config.KeyMark, config.Title)

	// update
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
	// del cache
	DelCache(form.Key, form.Title)

	// update db
	var err error
	switch form.Key {
	case common.WebsiteConfigKey:
		_, err = UpdateWebsite(form.Value)
		break
	case common.CorsConfigKey:
		_, err = UpdateCors(form.Value)
		break
	case common.StorageConfigKey:
		_, err = UpdateStorage(form.Value)
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

func DelCache(key, title string) {
	switch key {
	case common.WebsiteConfigKey:
		cache.DelWebsite()
		break
	case common.CorsConfigKey:
		cache.DelCors()
		break
	case common.ReserveHostConfigKey:
		cache.DelReverseHost(title)
		break
	default:
	}
}

func UpdateWebsite(value string) (*model.Website, error) {
	defaultWebsite := &model.Website{}
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultWebsite)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist, err := m.GetOneByKeyKeyMark(common.WebsiteConfigKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	} else {
		m.KeyMark = common.WebsiteConfigKey
		m.Value = value
		m.Title = common.WebsiteConfigKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultWebsite, nil
}

func UpdateCors(value string) (*model.Cors, error) {
	defaultCors := &model.Cors{}
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultCors)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist, err := m.GetOneByKeyKeyMark(common.CorsConfigKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	} else {
		m.KeyMark = common.CorsConfigKey
		m.Value = value
		m.Title = common.CorsConfigKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultCors, nil
}

func UpdateStorage(value string) (*model.Storage, error) {
	defaultStorage := &model.Storage{}
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultStorage)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist, err := m.GetOneByKeyKeyMark(common.StorageConfigKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	} else {
		m.KeyMark = common.StorageConfigKey
		m.Value = value
		m.Title = common.StorageConfigKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultStorage, nil
}
