package serverService

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"socialbot/internal/web/model"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/vultr"
)

func UpdateListRobotServerInfo() error {
	list := model.RobotServerList{}
	err := list.GetList()
	if err!= nil {
		return err
	}
	for _,rs := range list{
		err = UpdateRobotServerInfo(rs.Id, rs.ApiKey,rs.Subid)
		if err != nil {
			wblogger.Log.Error(err)
			continue
		}
	}
	return nil
}

func UpdateRobotServerInfo(id int, apiKey, subid string) error {
	server,err := vultr.GetServerInfo(apiKey, subid)
	if err != nil {
		return err
	}
	json,err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(server)
	if err != nil {
		return errors.Wrap(err, "json decode error")
	}
	m := model.RobotServer{
		FullServerInfo:string(json),
	}
	_, err = m.UpdateColsById(id, "full_server_info")
	if err != nil {
		return err
	}
	return nil
}
