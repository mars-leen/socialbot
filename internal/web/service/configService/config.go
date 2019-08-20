package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
)

var (
	defaultWebsite = &model.Website{
		HostName:"century",
	}
)

func GetWebsite() (*model.Website, error){
	f, r := cache.GetWebsite()
	if f {
		return r, nil
	}
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(common.WebsiteConfigKey)
	if err != nil {
		return nil, err
	}
	if !isExist {
		cache.SetWebsite(defaultWebsite)
		return defaultWebsite, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultWebsite)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}

	cache.SetWebsite(defaultWebsite)
	return defaultWebsite, nil
}


