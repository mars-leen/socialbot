package configService

import "socialbot/internal/web/common"

var (
	storage = &Storage{
		Source: common.DefaultStorageSource,
		ServeHost: "https://localhost:80",
		UploadLocalPath:  "",
		UploadLocalServePath: "/storage/",
	}
)

type Storage struct {
	Source           int `json:"Source"`
	ServeHost string `json:"ServeHost"`
	UploadLocalPath string `json:"UploadLocalPath"`
	UploadLocalServePath  string `json:"UploadLocalServePath"`
}

func GetStorage() (*Storage, error) {
	return storage, nil
}

func GetUploadFullUrl(base string) string {
	return storage.ServeHost + storage.UploadLocalServePath + base
}
