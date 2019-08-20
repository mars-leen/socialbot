package mediaLogic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"socialbot/internal/web/cache"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/mediaService"
	"socialbot/internal/web/service/tagService"
	"socialbot/internal/web/service/userService"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/utils"
	"strconv"
	"strings"
	"time"
)

//detail
func Detail(c *gin.Context,  uri int64) common.Result {
	media := model.Media{}
	isExist, err := media.GetOneByUri(uri)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	// add view num
	go func() {
		_, err := media.IncrById(media.Id, "view_num", nil)
		if err != nil {
			wblogger.Log.Error(err)
		}
	}()

	uid, err := userService.GetTokenUserId(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// get media source
	medias, err := mediaService.GetMediaSourceList(media.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// get media tags
	mediaTagList := &model.MediaTagList{}
	err = mediaTagList.GetMediaTagListByMidList([]int64{media.Id})
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	mediaTagListMap,err := GetMediaTagMidMap(mediaTagList)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	tagsValue, ok := mediaTagListMap[media.Id]
	if !ok {
		tagsValue = []model.ConTag{}
	}

	// isLike
	isLike, err := GetIsLike(media.Id, uid)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// commission product
	comProduct := model.ConComProduct{}
	if media.MediaType == common.MediaTypePromotionProduct {
		cp := model.CommissionProduct{}
		_,err := cp.GetOneByMid(media.Id)
		if err != nil {
			wblogger.Log.Error(err)
			return common.SystemError
		}
		comProduct.Link = cp.PromoteLink
		comProduct.CutOff = cp.CutOff
		comProduct.OriginPrice = cp.OriginPrice
		comProduct.NowPrice = cp.NowPrice
		comProduct.TotalStar = cp.TotalStar
		comProduct.NowStar = cp.NowStar
		comProduct.Reviews = cp.Reviews
	}

	// add view num
	return common.SUCCESS(model.ConMedia{
		Uri:         strconv.FormatInt(media.Uri, 10),
		Title:       media.Title,
		Medias:      medias,
		LikeNum:     media.LikeNum,
		ViewNum:     media.ViewNum,
		CommentNum:  media.CommentNum,
		MediaType:   media.MediaType,
		PublishAt:   media.PublishAt,
		Tags:        tagsValue,
		IsLike:      isLike,
		ComProduct:  comProduct,
	})
}


// list
func ListByRecommend(lastId int64, limit int) common.Result {
	m := model.MediaList{}
	err := m.GetRecommendList(lastId, limit)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	l := len(m)
	if l == 0 {
		return common.SUCCESSARR(nil)
	}

	conMediaList := make([]model.ConMedia, l)
	for i, value := range m {
		conMediaList[i] = model.ConMedia{
			Uri:        strconv.FormatInt(value.Uri, 10),
			Title:      value.Title,
			Cover:      configService.GetUploadFullUrl(value.Cover),
			MediaMum:   value.MediaNum,
			LikeNum:    value.LikeNum,
			ViewNum:    value.ViewNum,
			CommentNum: value.CommentNum,
			MediaType:  value.MediaType,
			PublishAt:  value.PublishAt,
			LastId:     strconv.FormatInt(value.Uri, 10),
		}
	}
	return common.SUCCESSARR(conMediaList)
}

func GetMediaTagMidMap(mtl *model.MediaTagList) (map[int64][]model.ConTag, error) {
	l := len(*mtl)
	mediaTagList := make(map[int64][]model.ConTag, l)

	if l == 0 {
		return mediaTagList, nil
	}
	tagsAll, err := tagService.GetTagsMapWithId()
	if err != nil {
		return mediaTagList, err
	}

	for _, value := range *mtl {
		if  tagValue, ok := tagsAll[value.Tid]; ok {
			mediaTagList[value.Mid] = append(mediaTagList[value.Mid], model.ConTag{
				Id:   tagValue.Id,
				Title: tagValue.Title,
				ShortName:tagValue.ShortName,
			})
		}
	}
	return mediaTagList, nil
}

func GetIsLike(mid, uid int64) (bool, error) {
	if uid == 0 {
		return false, nil
	}
	userLikeMedia := model.UserLikeMedia{}
	has, err := userLikeMedia.ExistOneByMidUid(mid, uid)
	if err != nil {
		return false, err
	}
	return has, nil
}

func ListByCategory(lastId int64, category, sort int) common.Result {
	ml :=model.MediaList{}
	var err error
	if sort == common.SortNew {
		err = ml.GetListNewestByCid(lastId, category)
	} else {
		err = ml.GetHotByCid(int(lastId), category)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	l := len(ml)
	if l == 0 {
		return common.SUCCESSARR(nil)
	}

	conMediaList := make([]model.ConMedia, l)
	lastSetId := strconv.FormatInt(lastId+1, 10)
	for i, value := range ml {
		if sort == common.SortNew {
			lastSetId = strconv.FormatInt(value.PublishAt, 10)
		}
		conMediaList[i] = model.ConMedia{
			Uri:        strconv.FormatInt(value.Uri, 10),
			Title:      value.Title,
			Cover:      configService.GetUploadFullUrl(value.Cover),
			MediaMum:   value.MediaNum,
			LikeNum:    value.LikeNum,
			ViewNum:    value.ViewNum,
			CommentNum: value.CommentNum,
			MediaType:  value.MediaType,
			PublishAt:  value.PublishAt,
			LastId:     lastSetId,
		}
	}
	return common.SUCCESSARR(conMediaList)
}

func Like(c *gin.Context, uri int64, delStatus int) common.Result {
	user, err := userService.MustGetTokenUser(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	mid, err := mediaService.GetMediaIdByUri(uri)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if mid == 0 {
		return common.DataIsNotExist
	}

	lock := cache.LockUserLikeMedia(user.Id, mid)
	if lock {
		return common.RepeatOperation
	}

	likeModel := model.UserLikeMedia{}
	exist, err := likeModel.GetOneByMidUid(mid, user.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// update status
	if exist {
		if likeModel.IsDel == delStatus {
			return common.RepeatOperation
		}
		err = UpdateLike(likeModel, delStatus)

		if err != nil {
			wblogger.Log.Error(err)
			return common.SystemError
		}
		return common.SUCCESS(nil)
	}

	// inset one
	err = InsertLike(mid, user.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	return common.SUCCESS(nil)
}

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
	var cover string
	for _,v := range mediaSourceList{
		if v.SourceType== common.SourceTypeImage {
			cover = v.Url
			break
		}
	}
	if cover == "" {
		return common.ParamError.Errorf("%s", "first img of uploaded file will be cover")
	}


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
		MediaNum:    l,
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
		Mid:         media.Id,
		CutOff:      form.CutOff,
		NowPrice:    form.NowPrice,
		OriginPrice: form.OriginPrice,
		TotalStar:   form.TotalStar,
		NowStar:     form.NowStar,
		Reviews:     form.Reviews,
		DetailLink:  form.DetailLink,
		PromoteLink: form.PromoteLink,
	}
	_, err = product.Insert(session)
	if err != nil {
		_ = session.Rollback()
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// update media source
	mediaSource := model.MediaSource{
		Mid: media.Id,
		Cid: form.Cid,
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
		Cid:         form.Cid,
		MediaNum:  len(mediaArr),
		MediaStatus: common.MediaStatusPublished,
		MediaType:   common.MediaTypePromotionProduct,
	}
	if form.Recommend {
		media.Recommend = 1
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
			Cid:        form.Cid,
			SourceType: value.FileType,
			SourceFrom: value.Source,
			Url:        value.FileName,
		}
		_, err := ms.Insert(session)
		if err != nil {
			wblogger.Log.Error(err)
			_ = session.Rollback()
			return common.SystemError
		}
	}

	ci := model.CrawlerItem{ItemStatus:common.CrawlerItemPublished}
	_, err =ci.UpdateColsById(form.CrawlerItemId, session, "item_status")
	if err != nil {
		wblogger.Log.Error(err)
		_ = session.Rollback()
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

func DecodeMedias(medias string) ([]*model.SubmitMediaForm, error) {
	ml := make([]*model.SubmitMediaForm, 0)
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(medias), &ml)
	if err != nil {
		return nil, errors.Wrap(err, "json decode medias failed")
	}
	return ml, nil
}

func DownLoadMedias(medias []*model.SubmitMediaForm) error {
	storage,err := configService.GetStorage()
	if err != nil {
		wblogger.Log.Error(err)
		return err
	}

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

func InsertLike(mid, uid int64) error {
	session := orm.SocialBotOrm.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	// add media num
	media := model.Media{}
	_, err = media.IncrById(mid, "like_num", session)
	if err != nil {
		_ = session.Rollback()
		return err
	}

	// insert media user relation
	likeModel := model.UserLikeMedia{}
	likeModel.Mid = mid
	likeModel.Uid = uid
	_, err = likeModel.Insert(session)
	if err != nil {
		_ = session.Rollback()
		return err
	}

	_ = session.Commit()
	return nil
}

func UpdateLike(likeModel model.UserLikeMedia, delStatus int) error {
	session := orm.SocialBotOrm.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	likeModel.IsDel = delStatus
	_, err = likeModel.UpDelById(likeModel.Id, session)
	if err != nil {
		_ = session.Rollback()
		return err
	}

	media := model.Media{}
	if delStatus == 1 {
		_, err = media.DecrById(likeModel.Mid, "like_num", session)
	} else {
		_, err = media.IncrById(likeModel.Mid, "like_num", session)
	}

	if err != nil {
		_ = session.Rollback()
		return err
	}
	_ = session.Commit()
	return nil
}
