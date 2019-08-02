package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/setting"
	"socialbot/pkg/snowflake"
	"time"
)

type User struct {
	Id         int64
	Uri        int64
	Avatar     string
	Email      string
	Identity   int
	Intro      string
	Nickname   string
	Password   string
	RegisterAt int64
	RegisterIp string
	UserStatus int
	UpdateAt   int64
	CreateAt   int64
	IsDel      int
}

type ConUser struct {
	Uri      int64  `json:"userUri"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

func (u *User) Insert() (rs int64, err error) {
	u.CreateAt = time.Now().UnixNano()
	u.Uri, err = snowflake.GetIntUUID(setting.NodeId)
	if err != nil {
		return 0, errors.Wrap(err, " snowflake.GetIntUUID failed")
	}
	rs, err = orm.SocialBotOrm.Insert(u)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (u *User) UpdateById(id int64) (rs int64, err error) {
	u.UpdateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Where("id = ? ", id).Update(u)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("updateById(%d) failed", id))
	}
	return rs, nil
}

func (u *User) UpdateColsById(id int64, cols ...string) (rs int64, err error) {
	u.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(u)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (u *User) GetOneById(id int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Get(u)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (u *User) GetColsOneById(id int64, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Cols(cols...).Get(u)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (u *User) GetOneByEmail(email string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("email=?", email).Get(u)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneByEmail(%s) failed", email))
	}
	return rs, nil
}

func (u *User) GetOneByUri(uri int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("uri=?", uri).Get(u)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneByUri(%s) cols(%s) failed", uri))
	}
	return rs, nil
}
