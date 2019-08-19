package robotServerLogic

import (
	jsoniter "github.com/json-iterator/go"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/vultr"
)

func Add(configId int, name string) common.Result {
	sc, err := configService.GetServersConfigById(configId)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	server, err := vultr.CreateServer(sc.ApiKey, name, sc.RegionID, sc.PlanId, sc.OsId, sc.ScriptName, sc.ScriptType,sc.ScriptContent)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError.Errorf("%v", err)
	}

	json, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(server)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	s := model.RobotServer{
		ApiKey:sc.ApiKey,
		Subid:server.ID,
		FullServerInfo:string(json),
	}
	_,err = s.Insert()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

func Delete(id int) common.Result {
	r := model.RobotServer{}
	isExist, err := r.GetOneById(id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	err = vultr.DeleteServer(r.ApiKey, r.Subid)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError.Errorf("%v",err)
	}

	r.IsDel = 1
	_,err = r.UpdateColsById(id, "is_del")
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESS(nil)
}

func List() common.Result {
	list := model.RobotServerList{}
	err := list.GetList()
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	return common.SUCCESSARR(list)
}