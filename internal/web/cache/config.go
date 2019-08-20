package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"socialbot/internal/web/model"
	"time"
)

const (
	WebsiteKey           = "config_website"    // config_website
	ReverseHostKey       = "config_reverse_%s" // config_host
	StorageKey           = "config_storage"    // config_storage
	CorsKey              = "config_cors"       // config_cors
	DefaultReverseExpire = 3 * 24 * time.Hour
)

var Config = cache.New(7*24*time.Hour, 7*24*time.Hour)

//website
func GetWebsite() (bool, *model.Website) {
	val, found := Config.Get(WebsiteKey)
	if !found {
		return found, nil
	}
	return found, val.(*model.Website)
}

func SetWebsite(value *model.Website) {
	Config.Set(WebsiteKey, value, DefaultReverseExpire)
}

func DelWebsite() {
	Config.Delete(WebsiteKey)
}

// storage
func GetStorage() (bool, *model.Storage) {
	val, found := Config.Get(StorageKey)
	if !found {
		return found, nil
	}
	return found, val.(*model.Storage)
}

func SetStorage(value *model.Storage) {
	Config.Set(StorageKey, value, DefaultReverseExpire)
}

func DelStorage() {
	Config.Delete(StorageKey)
}

// reverse host
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

// cors
func GetCors() (bool, *model.Cors) {
	val, found := Config.Get(CorsKey)
	if !found {
		return found, nil
	}
	return found, val.(*model.Cors)
}

func SetCors(value *model.Cors) {
	Config.Set(CorsKey, value, DefaultReverseExpire)
}

func DelCors() {
	Config.Delete(CorsKey)
}
