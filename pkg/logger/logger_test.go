package logger

import (
	"github.com/pkg/errors"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	l,err := NewFileLogger(Info, "./log/info.log", []int{Info,Error})
	if err != nil {
		panic(err)
	}

	l.Info(111, "aaa")
	l.Error(ER())
}

func ER() error {
	return  errors.Wrap(errors.New("aa"), "bbb")
}