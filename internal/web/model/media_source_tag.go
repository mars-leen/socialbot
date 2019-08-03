package model

import (
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type MediaSourceTag struct {
	Id             int64
	Cid            int
	Msid           int64
	Mid            int64
	Tid            int
	UpdateAt       int64
	CreateAt       int64
	IsDel          int
}

type MediaSourceTagList []MediaSource

func (m *MediaSourceTag) Insert(session *xorm.Session) (rs int64, err error) {
	m.CreateAt = time.Now().UnixNano()
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
