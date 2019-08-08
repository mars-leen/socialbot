package galleryLogic

import (
	"github.com/pkg/errors"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/mediaService"
	"socialbot/internal/web/service/tagService"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/utils"
	"strings"
)

func List(lastId int64, sort , sourceType int) common.Result {
	var err error
	m := model.MediaSourceList{}
	if sort == common.SortDesc {
		err = m.GetListDESC(lastId, sort, sourceType, common.GalleryPage)
	}else{
		err = m.GetListASC(lastId, sort, sourceType, common.GalleryPage)
	}
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// id list
	idList := make([]int64, len(m))
	for i,v := range m {
		idList[i] = v.Id
	}

	// get media source tag List by Id
	ms := model.MediaSourceTagList{}
	err = ms.GetListByMsidList(idList)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// get tag map
	mapTagList,err := mediaService.GetMapWithMediaSourceTag(ms)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	result := make([]model.ConMediaSourceWithTags, len(m))

	for i,v := range m{
		tags,ok := mapTagList[v.Id]
		if !ok {
			tags = []*model.ConTag{}
		}
		result[i] = model.ConMediaSourceWithTags{
			Id:v.Id,
			Url: configService.GetUploadFullUrl(v.Url),
			Uri:v.Uri,
			Cid:v.Cid,
			SourceFrom: v.SourceFrom,
			SourceType: v.SourceType,
			Tags:tags,
		}
	}
	return common.SUCCESSARR(result)
}

func AddTags(msid int64, cid int, tags string ) common.Result {
	tagList := strings.Split(utils.Trim(tags,",",","), ",")
	if len(tagList) == 0 {
		return common.TagsIsEmptyError
	}
	ms := model.MediaSource{}
	isExist, err := ms.GetOneById(msid)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !isExist {
		return common.DataIsNotExist
	}

	session := orm.SocialBotOrm.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		wblogger.Log.Error(errors.Wrap(err, "session.Begin failed"))
		return common.SystemError
	}

	err = tagService.InsertMediaSourceTag(tagList, ms.Mid, msid, ms.Cid, nil)
	if err != nil {
		wblogger.Log.Error(err)
		_ =session.Rollback()
		return common.SystemError
	}
	ms.Cid =cid
	_,err = ms.UpdateByCols(session, msid, "cid")
	if err != nil {
		wblogger.Log.Error(err)
		_ =session.Rollback()
		return common.SystemError
	}

	_ = session.Commit()
	return common.SUCCESS(nil)
}
