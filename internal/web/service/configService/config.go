package configService

import (
	"socialbot/internal/web/common"
)

var (
	host       = ""
	uploadHost = ""
	cors       = &Cors{
		Enable:           true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization", "x-requested-with"},
		AllowCredentials: true,
		MaxAge:           7200,
	}

	storage = &Storage{
		Source: common.DefaultStorageSource,
		ServeHost: "https://localhost:80",
		UploadLocalPath:  "",
		UploadLocalServePath: "/storage/",
	}
)

type Cors struct {
	Enable           bool     `json:"Enable"`
	AllowOrigins     []string `json:"AllowOrigins"`
	AllowMethods     []string `json:"AllowMethods"`
	AllowHeaders     []string `json:"AllowHeaders"`
	AllowCredentials bool     `json:"AllowCredentials"`
	MaxAge           int      `json:"MaxAge"`
}

type Storage struct {
	Source           int `json:"Source"`
	ServeHost string `json:"ServeHost"`
	UploadLocalPath string `json:"UploadLocalPath"`
	UploadLocalServePath  string `json:"UploadLocalServePath"`
}


func GetCors() *Cors {
	return cors
}

func GetStorage() (*Storage, error) {
	return storage, nil
}

func GetHostName() string {
	return host
}

func SetUploadHost(host string) {
	uploadHost = host
}

func GetUploadHost() string {
	return uploadHost
}

func GetUploadFullUrl(base string) string {
	return storage.ServeHost + storage.UploadLocalServePath + base
}
