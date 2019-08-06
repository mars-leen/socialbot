package vultr

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/pkg/vultr/lib"
	"strconv"
)

func GetClient(apiKey string) *lib.Client {
	return lib.NewClient(apiKey, nil)
}

func CreateScript(apiKey, name, content, stype string) (lib.StartupScript, error) {
	return GetClient(apiKey).CreateStartupScript(name, content, stype)
}

func UpdateOrCreateScriptByName(c *lib.Client, name, content, stype string) (s lib.StartupScript, err error) {
	list, err := c.GetStartupScripts()
	if err != nil {
		return s, err
	}
	var script lib.StartupScript
	isExist := false
	for _, v := range list {
		if v.Name == name && v.Type == stype {
			isExist = true
			script.ID = v.ID
			script.Name = v.Name
			script.Content = content
			break
		}
	}

	if isExist {
		err := c.UpdateStartupScript(script)
		if err != nil {
			return s, errors.Wrap(err, fmt.Sprintf("apikey %s find %s, update error",name))
		}
		return script, nil
	}

	s, err = c.CreateStartupScript(name, content, stype)
	if err != nil {
		return s, errors.Wrap(err, fmt.Sprintf("apikey %s find %s, create error", name))
	}
	return s, nil
}

func CreateServer(apiKey, name string, regionID, planID, osID int, scriptName, scriptType, scriptContent string) (*lib.Server, error) {
	c := GetClient(apiKey)
	script,err := UpdateOrCreateScriptByName(c, scriptName, scriptContent, scriptType)
	if err != nil {
		return nil,err
	}

	scriptId,err := strconv.Atoi(script.ID)
	if err != nil {
		return nil,errors.Wrap(err, "strconv.Atoi(script.ID) failed")
	}
	server, err := c.CreateServer(name, regionID, planID, osID, &lib.ServerOptions{
		Script: scriptId,
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("CreateServer faild, apikey %s name %s regionID %d  planID %d osID %d", apiKey, name, regionID, planID, osID))
	}
	return &server, nil
}

func GetServerInfo(apiKey, id string) (*lib.Server, error) {
	s,err := GetClient(apiKey).GetServer(id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(" apiKey %s, GetClient(apiKey).GetServer %s  faild", apiKey, id))
	}
	return &s, nil
}

func DeleteServer(apiKey,id string) error {
	err := GetClient(apiKey).DeleteServer(id)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("DeleteServer failed, apiKey %s ,id %s", apiKey, id))
	}
	return nil
}
