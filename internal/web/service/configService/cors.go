package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
)

var (
	defaultCors       = &model.Cors{
		Enable:           false,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization", "x-requested-with"},
		AllowCredentials: true,
		MaxAge:           7200,
	}
)

func GetCors() (*model.Cors, error){
	f, r := cache.GetCors()
	if f {
		return r, nil
	}

	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(common.CorsConfigKey)
	if err != nil {
		return nil, err
	}
	if !isExist {
		cache.SetCors(defaultCors)
		return defaultCors, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultCors)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	cache.SetCors(defaultCors)
	return defaultCors, nil
}

