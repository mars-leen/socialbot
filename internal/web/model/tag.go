package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type Tag struct {
	Id          int
	Cid         int
	Title       string
	BoardName   string
	Description string
	ShortName   string
	UpdateAt    int64
	CreateAt    int64
	IsDel       int
}

type ConTag struct {
	Id          int
	Cid         int
	Title       string
	BoardName   string
	Description string
	ShortName   string
}

type TagList []Tag

func (t *Tag) Insert() (rs int64, err error) {
	t.UpdateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(t)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (t *Tag) UpdateById(id int) (rs int64, err error) {
	rs, err = orm.SocialBotOrm.Where("id = ? ", id).Update(t)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("updateById(%d) failed", id))
	}
	return rs, nil
}

func (t *Tag) UpdateColsById(id int, cols ...string) (rs int64, err error) {
	t.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(t)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (t *Tag) GetOneById(id int) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Get(t)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (t *Tag) GetColsOneById(id int, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Cols(cols...).Get(t)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}


func (tl *TagList) GetListByCid(cid int) (err error) {
	err = orm.SocialBotOrm.Where("cid=? ", cid).Where("is_del=?", 0).Find(tl)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetListByCid(%d) failed", cid))
	}
	return nil
}

func (tl *TagList) GetList() (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Find(tl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}