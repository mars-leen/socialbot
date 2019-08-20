package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
)

func GetReverseHost(host string) (*model.ReserveHost, error) {
	found,r := cache.GetReverseHost(host)
	if found {
		return r, nil
	}
	// not found
	config := model.Config{}
	isExist,err := config.GetOneByKeAndTitle(common.ReserveHostConfigKey, host)
	if err != nil {
		return  nil, err
	}
	if !isExist {
		// 设置空值 防止 穿透
		cache.SetReverseHost(host, r)
		return  nil, nil
	}

	s := &model.ReserveHost{}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(config.Value), s)
	if err != nil {
		return  nil, errors.Wrap(err, "json decode error")
	}
	r = s
	cache.SetReverseHost(host, r)
	return r, nil
}