package vultr

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/pkg/vultr/lib"
)

func GetClient(apiKey string) *lib.Client {
	return lib.NewClient(apiKey, nil)
}

func CreateScript(apiKey, name, content, stype string) (lib.StartupScript, error) {
	return GetClient(apiKey).CreateStartupScript(name, content, stype)
}

func UpdateOrCreateScriptByName(apiKey, name, content, stype string) error {
	c := GetClient(apiKey)
	list,err := c.GetStartupScripts()
	if err != nil {
		return err
	}

	var script lib.StartupScript
	isExist := false
	for _,v := range list {
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
			return  errors.Wrap(err, fmt.Sprintf("apikey %s find %s, update error", apiKey, name))
		}
		return nil
	}

	_, err = c.CreateStartupScript(name, content, stype)
	if err != nil {
		return  errors.Wrap(err, fmt.Sprintf("apikey %s find %s, create error", apiKey, name))
	}
	return nil
}

func CreateServer(apiKey string) {
}