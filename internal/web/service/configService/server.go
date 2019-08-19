package configService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/model"
)

type Server struct {
	ScriptName  string
	ScriptType  string
	ScriptContent string
	ApiKey    string
	RegionID  int
	PlanId    int
	OsId      int
}

func GetServersConfigById(id int) (*Server, error) {
	config := model.Config{}
	_,err := config.GetOneById(id)
	if err != nil {
		return  nil, err
	}
	s := &Server{}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(config.Value), s)
	if err != nil {
		return  nil, errors.Wrap(err, "json decode error")
	}
	return s, err
}
