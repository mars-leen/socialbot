package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/model"
)

const (
	WebsiteKey = "website"
	ServerKey= "server"
	CorsKey = "cors"
	StorageKey = "storage"
	ReserveHostKey = "reserve_host"
)

var (
	defaultWebsite = &Website{
		HostName:"century",
	}
)

type Website struct {
	HostName   string
}

func GetHostName() string {
	return defaultWebsite.HostName
}

func GetWebsite() (*Website, error){
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(WebsiteKey)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return defaultWebsite, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultWebsite)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	return defaultWebsite, nil
}

func UpdateWebsite(value string) (*Website, error){
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultWebsite)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(WebsiteKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	}else{
		m.KeyMark = WebsiteKey
		m.Value = value
		m.Title = WebsiteKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultWebsite, nil
}


