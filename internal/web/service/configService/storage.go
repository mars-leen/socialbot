package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/setting"
	"strings"
)

var (
	defaultStorage = &Storage{
		Source:               common.DefaultStorageSource,
		ServeHost:            "",
		UploadLocalPath:      "",
		UploadLocalServePath: "/storage/",
	}
)

type Storage struct {
	Source               int    `json:"Source"`
	ServeHost            string `json:"ServeHost"`
	UploadLocalPath      string `json:"UploadLocalPath"`
	UploadLocalServePath string `json:"UploadLocalServePath"`
}


func GetStorage() (*Storage, error){
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(StorageKey)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return defaultStorage, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultStorage)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	return defaultStorage, nil
}

func UpdateStorage(value string) (*Storage, error){
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultStorage)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(StorageKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	}else{
		m.KeyMark = StorageKey
		m.Value = value
		m.Title = StorageKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultStorage, nil
}

func GetStorageUploadPath() string {
	if defaultStorage.UploadLocalPath == "" {
		return filepath.Join(setting.AppPath, "storage")
	}
	return defaultStorage.UploadLocalPath
}

func GetUploadFullUrl(base string) string {
	if base == "" {
		return base
	}
	return strings.ReplaceAll(defaultStorage.ServeHost + defaultStorage.UploadLocalServePath + base, `\`, `/`)
}
