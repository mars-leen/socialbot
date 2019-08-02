package wblogger

import (
	"path/filepath"
	"socialbot/internal/web/setting"
	"socialbot/pkg/logger"
)

var (
	Log *logger.Logger
	Request *logger.Logger
)

func LoadLogger() (err error) {
	Log,err = logger.NewFileLogger(logger.Info, filepath.Join(setting.AppPath, "log", "socialbot"), []int{logger.Info, logger.Error})
	if err != nil {
		return err
	}

	Request,err = logger.NewFileLogger(logger.Info, filepath.Join(setting.AppPath, "log", "request"), []int{logger.Info, logger.Error})
	if err != nil {
		return err
	}
	return nil
}
