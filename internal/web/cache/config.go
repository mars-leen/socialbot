package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"socialbot/internal/web/model"
	"time"
)
const (
	ReverseHostKey = "config_%s" // config_hos
	DefaultReverseExpire = 3*24*time.Hour
)

var Config = cache.New(7*24*time.Hour, 7*24*time.Hour)

func GetReverseHost(host string) (bool, *model.ReserveHost) {
	val, found := Config.Get(fmt.Sprintf(ReverseHostKey, host))
	if !found {
		return found, nil
	}
	return found, val.(*model.ReserveHost)
}

func SetReverseHost(host string, value *model.ReserveHost) {
	Config.Set(fmt.Sprintf(ReverseHostKey, host), value, DefaultReverseExpire)
}

func DelReverseHost(host string) {
	Config.Delete(fmt.Sprintf(ReverseHostKey, host))
}