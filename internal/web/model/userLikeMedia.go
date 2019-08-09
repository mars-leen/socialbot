package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type UserLikeMedia struct {
	Id          int64
	Mid         int64
	Uid         int64
	CreateAt    int64
	UpdateAt    int64
	IsDel       int
}

func (u *UserLikeMedia) Insert(session *xorm.Session) (rs int64, err error) {
	u.CreateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Insert(u)
		if err != nil {
			return 0, errors.Wrap(err, "Insert failed")
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Insert(u)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (u *UserLikeMedia) ExistByMidUid(mid,uid int64) (rs bool, err error) {
	rs,err = orm.SocialBotOrm.Where("mid=? and uid=?", mid, uid).Exist(u)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("ExistByMidUid(%d, %d) failed", mid, uid))
	}
	return rs, nil
}

func (u *UserLikeMedia) UpdateById(id int64) (rs int64, err error) {
	u.UpdateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Where("id = ? ", id).Update(u)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("updateById(%d) failed", id))
	}
	return rs, nil
}

func (u *UserLikeMedia) UpDelById(id int64, session *xorm.Session) (rs int64, err error) {
	u.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs , err = session.Cols("update_at","is_del").Where("id = ? ", id).Update(u)
		if err != nil {
			return 0, errors.Wrap(err, "UpDelById failed")
		}
		return rs, nil
	}

	rs , err = orm.SocialBotOrm.Cols("update_at","is_del").Where("id = ? ", id).Update(u)
	if err != nil {
		return 0, errors.Wrap(err, "UpDelById failed")
	}
	return rs, nil
}

func (u *UserLikeMedia) GetOneByMidUid(mid,uid int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("mid=? and uid=?", mid, uid).Get(u)
	if err != nil {
		return false , errors.Wrap(err, fmt.Sprintf("GetOneByMidUid(%d, %d) failed", mid, uid))
	}
	return rs, nil
}

func (u *UserLikeMedia) ExistOneByMidUid(mid,uid int64) (bool, error) {
	rs,err := orm.SocialBotOrm.Where("mid=? and uid=? and is_del=0", mid, uid).Exist(u)
	if err != nil {
		return false , errors.Wrap(err, fmt.Sprintf("ExistOneByMidUid(%d, %d) failed", mid, uid))
	}
	return rs, nil
}



