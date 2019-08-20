package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"path/filepath"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/setting"
	"strings"
)

var (
	defaultStorage = &model.Storage{
		Source:               common.DefaultStorageSource,
		ServeHost:            "",
		UploadLocalPath:      "",
		UploadLocalServePath: "/storage/",
	}
)

func GetStorage() (*model.Storage, error){
	found,r := cache.GetStorage()
	if found {
		return r, nil
	}
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(common.StorageConfigKey)
	if err != nil {
		return defaultStorage, err
	}
	if !isExist {
		cache.SetStorage(defaultStorage)
		return defaultStorage, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultStorage)
	if err != nil {
		return defaultStorage, errors.Wrap(err, "decode error")
	}
	cache.SetStorage(defaultStorage)
	return defaultStorage, nil
}

func GetStorageUploadPath() string {
	r,_ := GetStorage()
	if r.UploadLocalPath == "" {
		return filepath.Join(setting.AppPath, "storage")
	}
	return r.UploadLocalPath
}

func GetUploadFullUrl(base string) string {
	if base == "" {
		return base
	}
	r,_ := GetStorage()
	return strings.ReplaceAll(r.ServeHost + r.UploadLocalServePath + base, `\`, `/`)
}
