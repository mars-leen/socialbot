package configService

import "sync"

var servers  *ServerMap

type ServerMap struct {
	lock  *sync.RWMutex
	list map[string]Server
}

type Server struct {
	ScriptName  string
	ScriptType  string
	ScriptContent string
	ApiKey    string
	RegionID  int
	PlanId    int
	OsId      int
}

func SertServers(list []Server)  {
	servers.lock.Lock()
	for _,v := range list{
		servers.list[v.ApiKey] = v
	}
	servers.lock.Unlock()

	// db


}

func GetServers() *ServerMap {
	return servers
}
