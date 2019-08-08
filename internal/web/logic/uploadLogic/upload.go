package uploadLogic

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/mediaService"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/utils"
	"strconv"
)

func UploadSingle(c *gin.Context, fileKey string) common.Result {
	file, err := c.FormFile(fileKey)
	if err!= nil {
		wblogger.Log.Error(errors.Wrap(err, "upload failed"))
		return common.UploadFailed
	}
	ext,err := utils.GetFileExt(file)
	if err != nil {
		wblogger.Log.Error(err)
		return common.UploadFailed
	}

	fileType := utils.GetFileTypeByExt(ext)
	if fileType == "" {
		return common.InvalidFileFormat
	}

	// todo
	storage := configService.GetStorage()

	// storage path
	path := storage.UploadLocalPath
	fileName := filepath.Join(common.StorageMediaDir, utils.Md5(file.Filename) + ext)
	if path == ""{
		path = filepath.Join(setting.AppPath, "storage")
	}
	fullPath := filepath.Join(path, fileName)

	// create dir
	err = utils.SaveUploadedFile(file, fullPath)
	if err!= nil {
		wblogger.Log.Error(err)
		return common.UploadFailed
	}

	// save to database
	uri, err := mediaService.InsertTmpMediaSource(storage.Source, fileType, fileName)
	if err != nil {
		wblogger.Log.Error(err)
		return common.UploadFailed
	}

	return common.SUCCESS(map[string]string{
		"url" : configService.GetUploadFullUrl(fileName),
		"uri": strconv.FormatInt(uri, 10),
	})
}

