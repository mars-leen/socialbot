package configService

import (
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/setting"
)

var (
	storage = &Storage{
		Source:               common.DefaultStorageSource,
		ServeHost:            "https://localhost:80",
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

func GetStorage() *Storage {
	return storage
}

func GetStorageUploadPath() string {
	if storage.UploadLocalPath == "" {
		return filepath.Join(setting.AppPath, "storage")
	}
	return storage.UploadLocalPath
}

func GetUploadFullUrl(base string) string {
	return storage.ServeHost + storage.UploadLocalServePath + base
}
