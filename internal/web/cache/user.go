package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	LockUserLikeMediaKey = "lock_user_like_media_%d_%d"
)

var User = cache.New(time.Minute, time.Minute)

func LockUserLikeMedia(uid,mid int64) bool {
	key:= fmt.Sprintf(LockUserLikeMediaKey, uid,mid)
	_, found  :=User.Get(key)
	if found {
		return found
	}
	User.Set(key, 1, time.Second)
	return found
}

