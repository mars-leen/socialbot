package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type Crawler struct {
	Id            int
	CrawlerStatus int
	LastPage      string
	Name          string
	Script        string
	UpdateAt      int64
	CreateAt      int64
	IsDel         int
}

type CrawlerList []Crawler

func (c *Crawler) Insert() (rs int64, err error) {
	c.CreateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (c *Crawler) UpdateColsById(id int, cols ...string) (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (c *Crawler) GetOneById(id int) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (c *Crawler) GetColsOneById(id int, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Cols(cols...).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (cl *CrawlerList) GetList() (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}