package storageService

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"socialbot/internal/web/wblogger"
)

func SyncRemoveSigFile(source, path string) {
	if source == "local"{
		go func() {
			if path == ""{
				return
			}
			err := os.Remove(path)
			if err != nil {
				wblogger.Log.Error(errors.Wrap(err, fmt.Sprintf("remove file failed %s", path)))
			}
		}()
	}
}

