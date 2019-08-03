package model

import (
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/common"
	"socialbot/internal/web/orm"
	"time"
)

type MediaTag struct {
	Cid            int
	CreateAt       int64
	Id             int64
	IsDel          int
	MediaPublishAt int64
	MediaStatus    int
	Mid            int64
	Tid            int
	UpdateAt       int64
}

type MediaTagList []MediaTag

func (mt MediaTag) Insert(session *xorm.Session) (rs int64, err error) {
	mt.CreateAt = time.Now().UnixNano()
	if session!= nil {
		rs,err =session.Insert(mt)
		if err != nil {
			return rs, errors.Wrap(err, "Insert failed")
		}
		return rs, nil
	}
	rs,err = orm.SocialBotOrm.Insert(mt)
	if err != nil {
		return rs, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (mtl *MediaTagList) GetMediaTagList(lastId int64) error {
	where := "mid > ?"
	orderBy := "mid ASC"
	return orm.SocialBotOrm.Select("*").Where(where, lastId).Limit(common.DefaultPage).OrderBy(orderBy).Where("is_del=?", 0).Find(mtl)
}

func (mtl *MediaTagList) GetMediaTagListByMidList(midList []int64) error {
	return orm.SocialBotOrm.In("mid", midList).Where("is_del=?", 0).Find(mtl)
}

func (mtl *MediaTagList) GetMediaTagListByTidList(page, limit int, tidList []int) error {
	return orm.SocialBotOrm.In("tid", tidList).Limit(limit, page*limit).GroupBy("mid").Where("is_del=?", 0).Find(mtl)
}

func (mtl *MediaTagList) GetMediaTagListByTid(tid, limit int) error {
	return orm.SocialBotOrm.Where("tid = ? ", tid).Limit(limit).GroupBy("mid").Where("is_del=?", 0).Find(mtl)
}
