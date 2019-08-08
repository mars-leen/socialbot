package mediaLogic

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/mediaService"
	"socialbot/internal/web/service/tagService"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/utils"
	"strings"
	"time"
)

func AddCommissionProduct(form *model.CommissionProductForm) common.Result {
	// media
	mediaArr := strings.Split(utils.Trim(form.Medias, ",", ","), ",")
	l := len(mediaArr)
	if l == 0 {
		return common.UploadFileIsEmpty
	}
	mediaArrList := utils.Str2Int64(mediaArr)

	// get origin media source
	mediaSourceList, err := mediaService.GetMapUriMediaSource(mediaArrList)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if l != len(mediaSourceList) {
		return common.UploadFileNotFound
	}

	//  cover
	v, ok := mediaSourceList[mediaArrList[0]]
	if !ok {
		return common.UploadFileNotFound
	}
	cover := v.Url

	session := orm.SocialBotOrm.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		wblogger.Log.Error(errors.Wrap(err, "session.Begin failed"))
		return common.SystemError
	}

	// insert media
	media := model.Media{
		Title:       form.Title,
		Cover:       cover,
		MediaStatus: common.MediaStatusPublished,
		MediaType:   common.MediaTypePromotionProduct,
	}
	_, err = media.Insert(session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// commission product
	product := model.CommissionProduct{
		Mid:media.Id,
		CutOff:form.CutOff,
		NowPrice:form.NowPrice,
		OriginPrice:form.OriginPrice,
		TotalStar:form.TotalStar,
		NowStar:form.NowStar,
		Reviews:form.Reviews,
		DetailLink:form.DetailLink,
		PromoteLink:form.PromoteLink,
	}
	_,err = product.Insert(session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// update media source
	mediaSource := model.MediaSource{
		Mid: media.Id,
	}
	_, err = mediaSource.UpdateColsByUriList(mediaArrList, session, "mid")
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// tag
	err = tagService.InsertMediaTag(form.Tags, media.Id, media.PublishAt, media.MediaStatus, form.Cid, session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	_ = session.Commit()
	return common.SUCCESS(nil)
}

func AddSocialMediaFromCrawler(form *model.SocialProductForm) common.Result {
	// media
	mediaArr, err := DecodeMedias(form.Medias)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	l := len(mediaArr)
	if l == 0 {
		return common.UploadFileIsEmpty
	}

	// fetcher file
	err = DownLoadMedias(mediaArr)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// first image is cover todo video cover
	var cover string
	for _, ms := range mediaArr {
		if ms.FileType == common.SourceTypeImage {
			cover = ms.FileName
			break
		}
	}

	// insert to db
	session := orm.SocialBotOrm.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		wblogger.Log.Error(errors.Wrap(err, "session.Begin failed"))
		return common.SystemError
	}

	// insert media
	media := model.Media{
		Title:       form.Title,
		Cover:       cover,
		MediaStatus: common.MediaStatusPublished,
		MediaType:   common.MediaTypePromotionProduct,
	}
	_, err = media.Insert(session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// update media source
	for _, value := range mediaArr {
		ms := model.MediaSource{
			Mid:        media.Id,
			SourceType: value.FileType,
			SourceFrom: value.Source,
			Url:        value.FileName,
		}
		_, err := ms.Insert(session)
		if err != nil {
			wblogger.Log.Error(err)
			return common.SystemError
		}
	}

	// tag
	err = tagService.InsertMediaTag(form.Tags, media.Id, media.PublishAt, media.MediaStatus, form.Cid, session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	_ = session.Commit()
	return common.SUCCESS(nil)
}

func DecodeMedias(medias string) ([]*model.SubmitMediaForm, error) {
	ml := make([]*model.SubmitMediaForm, 0)
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(medias), &ml)
	if err != nil {
		return nil, errors.Wrap(err, "json decode medias failed")
	}
	return ml, nil
}

func DownLoadMedias(medias []*model.SubmitMediaForm) error {
	storage := configService.GetStorage()
	f := &http.Client{
		Timeout: 300 * time.Second,
	}

	path := storage.UploadLocalPath
	if path == "" {
		path = filepath.Join(setting.AppPath, "storage")
	}

	for _, m := range medias {
		fileType, filename, err := DownloadFile(f, m.URL, path, common.StorageMediaDir)
		if err != nil {
			return err
		}
		ft := common.SourceTypeImage
		if fileType == "VIDEO" {
			ft = common.SourceTypeVideo
		}
		m.FileName = filename
		m.FileType = ft
		m.Source = storage.Source
	}
	return nil
}

func DownloadFile(f *http.Client, requestUrl, storagePath, subPath string) (mediaType string, filename string, err error) {
	err = os.MkdirAll(filepath.Join(storagePath, subPath), os.ModePerm)
	if err != nil {
		return "", "", errors.Wrap(err, "mkdir all failed")
	}
	// Get the data
	/*c := http.Client{
		Timeout: 120 * time.Second,
	}*/

	/*if setting.Cfg.HTTP.DownloadProxy != "" {
		u, _ := url.Parse(setting.Cfg.HTTP.DownloadProxy)
		transport  := &http.Transport{
			Proxy: http.ProxyURL(u),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		c.Transport = transport
	}
	*/
	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = f.Get(requestUrl)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "Timeout") {
			break
		}
	}
	if err != nil {
		return "", "", errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the file
	ext, err := utils.GetFileExtFromResp(resp)
	if err != nil {
		return "", "", err
	}
	fileType := utils.GetFileTypeByExt(ext)
	filename = filepath.Join(subPath, utils.Md5(requestUrl)+ext)
	fullPath := filepath.Join(storagePath, filename)

	out, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fileType, filename, err
	}
	defer out.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fileType, filename, errors.Wrap(err, "io.copy error")
	}
	return fileType, filename, nil
}
