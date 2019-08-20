package cache

import (
	"fmt"
	"testing"
)

func TestSetReverseHost(t *testing.T) {
	SetReverseHost("aaa", nil)
	f, v := GetReverseHost("aaa")
	fmt.Println(f, v)
	DelReverseHost("aaa")
	f, v = GetReverseHost("aaa")
	fmt.Println(f, v)
}

