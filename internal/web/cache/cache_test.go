package cache

import (
	"fmt"
	"testing"
	"time"
)

func Test_SetLock(t *testing.T)  {
	LikeLock.Set("tt", 11, 2*time.Second)

	for i:=0;i<10;i++ {
		val, found  :=LikeLock.Get("tt")
		if found {
			fmt.Println(val.(int))
		}
		time.Sleep(time.Second)
	}
}

func TestSetReverseHost(t *testing.T) {
	SetReverseHost("aaa", nil)
	f, v := GetReverseHost("aaa")
	fmt.Println(f, v)
	DelReverseHost("aaa")
	f, v = GetReverseHost("aaa")
	fmt.Println(f, v)
}

