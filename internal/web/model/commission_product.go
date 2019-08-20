package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type CommissionProduct struct {
	Content     string
	CreateAt    int64
	DetailLink  string
	Id          int64
	IsDel       int
	Mid         int64
	PromoteLink string
	UpdateAt    int64
}

type ConComProduct struct {
	Link   string
}

func (c *CommissionProduct) Insert(session *xorm.Session) (rs int64, err error) {
	c.CreateAt = time.Now().UnixNano()
	if session != nil {
		rs,err = session.Insert(c)
		if err != nil {
			return rs, errors.Wrap(err, "Insert failed")
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return rs, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}


func (c *CommissionProduct) GetOneByMid(mid int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("mid=?", mid).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneByMid(%d) failed", mid))
	}
	return rs, nil
}