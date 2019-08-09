package uploadLogic

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/mediaService"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/snowflake"
	"socialbot/pkg/utils"
	"strconv"
)

func UploadSingle(c *gin.Context, fileKey string) common.Result {
	file, err := c.FormFile(fileKey)
	if err != nil {
		wblogger.Log.Error(errors.Wrap(err, "upload failed"))
		return common.UploadFailed
	}
	ext, err := utils.GetFileExt(file)
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
	fileName := filepath.Join(common.StorageMediaDir, utils.Md5(file.Filename)+ext)
	if path == "" {
		path = filepath.Join(setting.AppPath, "storage")
	}
	fullPath := filepath.Join(path, fileName)

	// create dir
	err = utils.SaveUploadedFile(file, fullPath)
	if err != nil {
		wblogger.Log.Error(err)
		return common.UploadFailed
	}

	// save to database
	uri, err := mediaService.InsertTmpMediaSource(storage.Source, fileType, fileName)
	if err != nil {
		wblogger.Log.Error(err)
		return common.UploadFailed
	}

	return common.SUCCESS(map[string]interface{}{
		"url": configService.GetUploadFullUrl(fileName),
		"uri": strconv.FormatInt(uri, 10),
		"fileType":fileType,
		"fileName":fileName,
		"Source": storage.Source,
	})

}

func UploadAvatar(c *gin.Context) (string, error) {
	avatar, err := c.FormFile("avatar")
	if err != nil && err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", errors.Wrap(err, "upload avatar failed")
	}

	exit, err := utils.GetFileExt(avatar)
	if err != nil {
		return "", err
	}

	Type := utils.GetFileTypeByExt(exit)
	if Type != "IMG" {
		return "", fmt.Errorf("inviald img")
	}

	randName, err := snowflake.GetStringUUID(setting.NodeId)
	if err != nil {
		return "", err
	}

	storagePath := configService.GetStorageUploadPath()
	fileName := filepath.Join(common.StorageAvatarDir, randName+exit)
	fullPath := filepath.Join(storagePath, fileName)
	err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "mkdir all failed")
	}

	//err = utils.SaveUploadedFile(avatar, fullPath)
	src, err := avatar.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, err := imaging.Decode(src)
	if err != nil {
		return "", err
	}
	resultImg := imaging.Fill(img, 128, 128, imaging.Center, imaging.Lanczos)
	err = imaging.Save(resultImg, fullPath)
	if err != nil {
		return "", errors.Wrap(err, "cut avatar failed")
	}
	return fileName, nil
}
