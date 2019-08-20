package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
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
	ItemStatus  int
	Title       string
	UpdateAt    int64
	CreateAt    int64
	IsDel       int
}


type CrawlerItemList []CrawlerItem

type ConCrawlerItem struct {
	Id          int64
	Crwid       int
	Medias      []CrawlerMedia
	Cover       string
	Description string
	Title       string
}

type CrawlerMedia struct {
	Show     string
	Download string
}


func (c *CrawlerItem) Insert() (rs int64, err error) {
	c.CreateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (c *CrawlerItem) UpdateColsById(id int64, session *xorm.Session, cols ...string) (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")

	if session != nil {
		rs, err = session.Cols(cols...).Where("id = ? ", id).Update(c)
		if err != nil {
			return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (c *CrawlerItem) GetOneById(id int64) (rs bool, err error) {
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

/*func (cl *CrawlerItemList) GetListByCrwid(crwid int, lastId int64, limit int) (err error) {
	err = orm.SocialBotOrm.Where("crwid=? ", crwid).Where("is_del=?", 0).Find(cl)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetListByCrwid(%d) failed", crwid))
	}
	return nil
}

func (cl *CrawlerItemList) GetList(lastId int64,limit int) (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Limit(limit).Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}*/

func (c *CrawlerItem) CountAll(crwid, status int) (int64, error) {
	where := "item_status = ? AND is_del = 0 "
	param := []interface{}{status}
	if crwid == 0 {
		where = "item_status = ? AND crwid= ? AND is_del = 0 "
		param = []interface{}{status, crwid}
	}
	count, err := orm.SocialBotOrm.Cols("id").Where(where, param...).Count(c)
	if err != nil {
		return 0, errors.Wrap(err, "get list failed")
	}
	return count, nil
}

func (cl *CrawlerItemList) GetListByMidList(idList []int64, status int) error {
	err := orm.SocialBotOrm.Omit("create_at", "status", "is_del").Where("status = ? AND is_del = 0", status).In("id", idList).Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}

func (cl *CrawlerItemList) GetRandList(crwid, status, limit int, ) error {
	where := "item_status = ? AND is_del = 0 "
	param := []interface{}{status}
	if crwid != 0 {
		where = "item_status = ? AND crwid= ? AND is_del = 0 "
		param = []interface{}{status, crwid}
	}

	err := orm.SocialBotOrm.Where(where, param...).OrderBy("RAND()").Limit(limit).Find(cl)

	if err != nil {
		return errors.Wrap(err, "get GeRandList failed")
	}
	return nil
}
