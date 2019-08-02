package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var LikeLock = cache.New(time.Minute, time.Minute)
var Config = cache.New(7*24*time.Hour, 7*24*time.Hour)