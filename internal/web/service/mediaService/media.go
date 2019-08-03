package mediaService

import (
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
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


