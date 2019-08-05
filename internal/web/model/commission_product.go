package model

import (
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type CommissionProduct struct {
	Content     string
	CreateAt    int64
	CutOff      int
	DetailLink  string
	Id          int64
	IsDel       int
	Mid         int64
	NowPrice    int
	NowStar     int
	OriginPrice int
	PromoteLink string
	Reviews     int
	TotalStar   int
	UpdateAt    int64
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