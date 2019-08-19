package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/model"
)

var (
	defaultCors       = &Cors{
		Enable:           true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization", "x-requested-with"},
		AllowCredentials: true,
		MaxAge:           7200,
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

func GetCors() (*Cors, error){
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(CorsKey)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return defaultCors, nil
	}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(m.Value), defaultCors)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	return defaultCors, nil
}

func UpdateCors(value string) (*Cors, error){
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(value), defaultCors)
	if err != nil {
		return nil, errors.Wrap(err, "decode error")
	}
	m := model.Config{}
	isExist,err := m.GetOneByKeyKeyMark(CorsKey)
	if err != nil {
		return nil, err
	}
	if isExist {
		m.Value = value
		_, err = m.UpdateColsById(m.Id, "value")
	}else{
		m.KeyMark = CorsKey
		m.Value = value
		m.Title = CorsKey
		_, err = m.Insert()
	}
	if err != nil {
		return nil, err
	}
	return defaultCors, nil
}