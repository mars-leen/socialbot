package tagService

import (
	"github.com/go-xorm/xorm"
	"socialbot/internal/web/model"
	"socialbot/pkg/utils"
	"strconv"
	"strings"
)

func GetTagsMapWithCid() (rs map[int][]model.ConTag, err error) {
	list := model.TagList{}
	err = list.GetList()
	if err != nil {
		return rs, err
	}

	l := len(list)
	if l == 0 {
		return map[int][]model.ConTag{}, nil
	}

	result := make(map[int][]model.ConTag, l)
	for _, t := range list {
		conTag := model.ConTag{
			Id:          t.Id,
			Title:       t.Title,
			ShortName:   t.ShortName,
			Description: t.Description,
			BoardName:   t.BoardName,
		}

		if _, exist := result[t.Cid]; !exist {
			result[t.Cid] = []model.ConTag{conTag}
		} else {
			result[t.Cid] = append(result[t.Cid], conTag)
		}
	}

	return result, nil
}

func GetTagsMapWithId() (rs map[int]model.ConTag, err error) {
	list := model.TagList{}
	err = list.GetList()
	if err != nil {
		return rs, err
	}

	l := len(list)
	if l == 0 {
		return map[int]model.ConTag{}, nil
	}

	result := make(map[int]model.ConTag, l)
	for _, t := range list {
		conTag := model.ConTag{
			Id:          t.Id,
			Title:       t.Title,
			ShortName:   t.ShortName,
			Description: t.Description,
			BoardName:   t.BoardName,
		}
		result[t.Id] = conTag
	}

	return result, nil
}

func InsertMediaTag(tagStr string, mediaId, publishTime int64, mediaStatus, cid int, session *xorm.Session) error {
	tagStr = utils.Trim(tagStr, ",", ",")
	if tagStr == "" {
		return nil
	}
	tagList := strings.Split(tagStr, ",")
	tags, err := GetTagsMapWithId()
	if err != nil {
		return err
	}
	m := model.MediaTag{}
	for _, t := range tagList {
		tId, _ := strconv.Atoi(t)
		if _, ok := tags[tId]; !ok {
			continue
		}
		m.Tid = tId
		m.Mid = mediaId
		m.Cid = cid
		m.MediaStatus = mediaStatus
		m.MediaPublishAt = publishTime
		_, err = m.Insert(session)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertMediaSourceTag(tagList []string, mediaId, MediaSourceId int64, cid int, session *xorm.Session) error {
	tags, err := GetTagsMapWithId()
	if err != nil {
		return err
	}
	for _, tt := range tagList {

		 t,_:= strconv.Atoi(tt)
		if _, ok := tags[t]; !ok {
			continue
		}
		m := model.MediaSourceTag{}
		m.Tid = t
		m.Mid = mediaId
		m.Msid = MediaSourceId
		m.Cid = cid
		_, err = m.Insert(session)
		if err != nil {
			return err
		}
	}
	return nil
}