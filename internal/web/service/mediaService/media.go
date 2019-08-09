package mediaService

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/tagService"
)

func InsertTmpMediaSource(source int,  mediaType, url string) (int64, error) {
	sourceType := common.SourceTypeImage
	if mediaType == "VIDEO"{
		sourceType = common.SourceTypeVideo
	}

	ms := model.MediaSource{
		SourceFrom:source,
		SourceType:sourceType,
		Url:url,
	}
	_,err := ms.Insert(nil)
	if err != nil {
		return 0, err
	}
	return ms.Uri, nil
}

func GetMapUriMediaSource(uriList []int64) (map[int64]model.MediaSource, error){
	list := model.MediaSourceList{}
	err :=list.GetListByUriList(uriList)
	if err!= nil {
		return nil, err
	}
	l := len(list)
	if l == 0 {
		return map[int64]model.MediaSource{}, nil
	}

	result := make(map[int64]model.MediaSource, l)

	for _,v := range list {
		result[v.Uri] = v
	}
	return result, nil
}

func GetMapWithMediaSourceTag(list model.MediaSourceTagList) (map[int64][]*model.ConTag, error) {
	allTag, err := tagService.GetTagsMapWithId()
	if err != nil {
		return nil, err
	}
	m := make(map[int64][]*model.ConTag)
	for _,ms := range list{

		tag,ok := allTag[ms.Tid]
		if !ok {
			tag = model.ConTag{}
		}
		if _, ok := m[ms.Msid];ok{
			m[ms.Msid] =append(m[ms.Msid], &tag)
		}else{
			m[ms.Msid] =[]*model.ConTag{&tag}
		}
	}
	return m, nil
}

func GetMediaIdByUri(uri int64) (int64, error) {
	// todo cache
	media := model.Media{}
	b, err := media.GetIdByUri(uri)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, nil
	}
	return media.Id, nil
}

func GetMediaIdList(ml *model.MediaList) []int64 {
	l := len(*ml)
	idList := make([]int64, l)
	if l == 0 {
		return idList
	}
	for i, value := range *ml {
		idList[i] = value.Id
	}
	return idList
}

func GetMediaSourceList(mid int64) ([]model.ConMediaSource, error) {
	msl := model.MediaSourceList{}
	err := msl.GetListByMid(mid)
	if err != nil {
		return nil, err
	}
	l:= len(msl)
	if len(msl) == 0 {
		return []model.ConMediaSource{}, nil
	}
	res := make([]model.ConMediaSource, l)
	for i,v :=  range msl {
		res[i] = model.ConMediaSource{
			Url: configService.GetUploadFullUrl(v.Url),
			SourceType: v.SourceType,
		}
	}
	return res, nil
}