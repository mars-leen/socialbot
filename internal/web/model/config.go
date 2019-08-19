package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type Config struct {
	CreateAt int64
	Id       int
	IsDel    int
	KeyMark      string
	Title    string
	UpdateAt int64
	Value    string
}

type ConfigList []Config

func (c *Config) Insert() (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (c *Config) UpdateById(id int) (rs int64, err error) {
	rs, err = orm.SocialBotOrm.Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("updateById(%d) failed", id))
	}
	return rs, nil
}

func (c *Config) UpdateColsById(id int, cols ...string) (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "update_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (c *Config) GetOneById(id int) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (c *Config) GetOneByKeyKeyMark(key string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("key_mark=?", key).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneByKeyKeyMark(%s) failed", key))
	}
	return rs, nil
}

func (c *Config) GetOneByKeAndTitle(key,title string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("key_mark=? AND title=?", key, title).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneByKeAndTitle(%s,%s) failed", key,title))
	}
	return rs, nil
}

func (c *Config) GetColsOneById(id int, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Cols(cols...).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (cl *ConfigList) GetList() (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).OrderBy("id desc").Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}

func (cl *ConfigList) GetListByKey(key string) (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Where("key_mark=?", key).OrderBy("id desc").Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}