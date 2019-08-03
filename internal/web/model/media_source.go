package model

import (
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
	Uri        int64
	IsDel      int
	Mid        int64
	SourceFrom int
	SourceType int
	UpdateAt   int64
	Url        string
}

type ConMediaSource struct {
	Id         int64
	Uri        int64
	Mid        int64
	SourceFrom int
	SourceType int
}

type MediaSourceList []MediaSource

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
	err := orm.SocialBotOrm.In("uri", uriList).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListByUriListMap failed")
	}
	return nil
}
