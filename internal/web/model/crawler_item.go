package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type CrawlerItem struct {
	Id          int64
	Crwid       int
	Content     string
	Cover       string
	Description string
	Title       string
	UpdateAt    int64
	CreateAt    int64
	IsDel       int
}

type CrawlerItemList []CrawlerItem

func (c *CrawlerItem) Insert() (rs int64, err error) {
	c.CreateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (c *CrawlerItem) UpdateColsById(id int, cols ...string) (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (c *CrawlerItem) GetOneById(id int) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (c *CrawlerItem) GetColsOneById(id int, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Cols(cols...).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (cl *CrawlerItemList) GetListByCrwid(crwid int) (err error) {
	err = orm.SocialBotOrm.Where("crwid=? ", crwid).Where("is_del=?", 0).Find(cl)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetListByCrwid(%d) failed", crwid))
	}
	return nil
}

func (cl *CrawlerItemList) GetList() (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}



