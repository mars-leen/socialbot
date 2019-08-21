package model

import (
	"fmt"
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

type MediaSourceTagList []MediaSourceTag

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

func (m *MediaSource) UpdateColsByIdList(idList []int64, session *xorm.Session, cols ...string) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Cols(cols...).In("id", idList).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "UpdateColsByIdList failed")
		}
		return rs, nil
	}
	rs, err = orm.SocialBotOrm.Cols(cols...).In("id", idList).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "UpdateColsByIdList failed")
	}
	return rs, nil
}

// list tag
func (ml *MediaSourceTagList) GetListByMsidList(msidList []int64) (err error) {
	err =  orm.SocialBotOrm.Omit("create_at","update_at" ).In("msid", msidList).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetListByMidList failed, list %v",msidList))
	}
	return  err
}


func (ml *MediaSourceTagList) GetAllListByMsid(msid int64) (err error) {
	err =  orm.SocialBotOrm.Omit("create_at","update_at" ).Where("msid=?",msid).Find(ml)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetAllListByMsid failed, list %v",msid))
	}
	return  err
}

func (m *MediaSourceTag) DeleteByMsid(msid int64, session *xorm.Session) (rs int64, err error) {
	if session != nil {
		rs, err = session.Where("msid=?", msid).Delete(m)
		if err != nil {
			return 0, errors.Wrap(err, "DeleteByMsid failed")
		}
		return rs, nil
	}
	rs, err = orm.SocialBotOrm.Where("msid=?", msid).Delete(m)
	if err != nil {
		return 0, errors.Wrap(err, "DeleteByMsid failed")
	}
	return rs, nil
}

