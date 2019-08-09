package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/setting"
	"socialbot/pkg/snowflake"
	"time"
)

type MediaSource struct {
	CreateAt   int64
	Id         int64
	Cid        int
	Uri        int64
	IsDel      int
	Mid        int64
	SourceFrom int
	SourceType int
	UpdateAt   int64
	Url        string
}

type ConMediaSource struct {
	Url        string
	SourceType int
}

type MediaSourceList []MediaSource

type ConMediaSourceWithTags struct {
	Id  int64
	Url string
	Uri int64
	Cid  int
	SourceFrom int
	SourceType int
	Tags  []*ConTag
}

func (m *MediaSource) GetOneById(id int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Get(m)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (m *MediaSource) Insert(session *xorm.Session) (rs int64, err error) {
	m.CreateAt = time.Now().UnixNano()
	m.Uri, err = snowflake.GetIntUUID(setting.NodeId)
	if err != nil {
		return 0, errors.Wrap(err, " snowflake.GetIntUUID failed")
	}

	if session != nil {
		rs, err = session.Insert(m)
		if err != nil {
			return 0, errors.Wrap(err, "Insert failed")
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Insert(m)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (m *MediaSource) UpdateColsByUriList(uriList []int64, session *xorm.Session, cols ...string) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Cols(cols...).In("uri", uriList).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "UpdateColsByUriList failed")
		}
		return rs, nil
	}
	rs, err = orm.SocialBotOrm.Cols(cols...).In("uri", uriList).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "UpdateColsByUriList failed")
	}
	return rs, nil
}

func (ml *MediaSourceList) GetListByUriList(uriList []int64) error {
	err := orm.SocialBotOrm.In("uri", uriList).Where("is_del=?", 0).OrderBy("id asc").Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListByUriList failed")
	}
	return nil
}

func (ml *MediaSourceList) GetListDESC(lastId int64, sort, sourceType, limit int) error {
	where := "id < ?"
	orderBy := "id desc"
	param := []interface{}{lastId}
	if lastId == 0 {
		where = "id > ?"
	}

	if sourceType != 0 {
		where = where + " AND source_type = ?"
		param = []interface{}{lastId, sourceType}
	}
	err := orm.SocialBotOrm.Where(where, param...).Limit(limit).OrderBy(orderBy).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListASC failed")
	}
	return nil
}

func (ml *MediaSourceList) GetListASC(lastId int64, sort, sourceType, limit int) error {
	where := "id > ? "
	orderBy := "id ASC"
	param := []interface{}{lastId}
	if sourceType != 0 {
		where = where + " AND source_type = ?"
		param = []interface{}{lastId, sourceType}
	}
	err := orm.SocialBotOrm.Where(where, param...).Limit(limit).OrderBy(orderBy).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListASC failed")
	}
	return nil
}

func (m *MediaSource) UpdateByCols(session *xorm.Session, id int64, cols ...string) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()

	if session != nil {
		rs,err= session.Cols(cols...).Where("id = ? ", id).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "UpdateByCols failed")
		}
		return rs, nil
	}
	rs,err= orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "UpdateByCols failed")
	}
	return rs, nil
}


func (ml *MediaSourceList) GetListByMid(mid int64) error {
	err := orm.SocialBotOrm.Where("mid=? and is_del=?", mid, 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListASC failed")
	}
	return nil
}