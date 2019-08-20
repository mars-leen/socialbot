package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	UriRlMidKey = "media_uri_%d"
	DefaultMediaExpire = 3*24*time.Hour
)

var Media = cache.New(time.Minute, time.Minute)

func GetMidByUri(uri int64) int64 {
	v, f :=Media.Get(fmt.Sprintf(UriRlMidKey, uri))
	if f {
		return v.(int64)
	}
	return 0
}

func SetMidByUri(uri, mid int64) {
	Config.Set(fmt.Sprintf(UriRlMidKey, uri), mid, DefaultMediaExpire)
}