package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var LikeLock = cache.New(time.Minute, time.Minute)


