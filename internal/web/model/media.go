package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"socialbot/internal/web/common"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/setting"
	"socialbot/pkg/snowflake"
	"time"
)

type Media struct {
	Cid         int
	CommentNum  int
	CreateAt    int64
	Id          int64
	IsDel       int
	LikeNum     int
	MediaNum    int
	MediaStatus int
	MediaType   int
	PublishAt   int64
	Title       string
	Cover       string
	Recommend   int
	Uid         int64
	UpdateAt    int64
	Uri         int64
	ViewNum     int
}

type ConMedia struct {
	Uri        string
	Cover      string
	Title      string
	MediaMum   int
	LikeNum    int
	ViewNum    int
	CommentNum int
	MediaType  int
	PublishAt  int64
	LastId     string
	Cid        int
	Recommend  int
	Tags       []ConTag
	IsLike     bool
	Medias     []ConMediaSource
	ComProduct ConComProduct
}

type MediaList []*Media

func (m *Media) Insert(session *xorm.Session) (rs int64, err error) {
	m.CreateAt = time.Now().UnixNano()
	if m.MediaStatus == common.MediaStatusPublished {
		m.PublishAt = m.CreateAt
	}
	m.Uri, err = snowflake.GetIntUUID(setting.NodeId)
	if err != nil {
		return rs, errors.Wrap(err, "snowflake.GetIntUUID failed")
	}
	if session != nil {
		rs, err = session.Insert(m)
		if err != nil {
			return rs, errors.Wrap(err, "Insert failed")
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Insert(m)
	if err != nil {
		return rs, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (m *Media) UpdateByCols( session *xorm.Session, id int64, cols ...string) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Cols(cols...).Where("id = ? ", id).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "UpdateByCols failed")
		}
		return rs, nil
	}

	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "UpdateByCols failed")
	}
	return rs, nil
}

func (m *Media) IncrById(id int64, col string, session *xorm.Session) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Cols(col).Where("id = ? ", id).Incr(col, 1).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "IncrById failed")
		}
		return rs, nil
	}
	rs, err = orm.SocialBotOrm.Cols(col).Where("id = ? ", id).Incr(col, 1).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "IncrById failed")
	}
	return rs, nil
}

func (m *Media) DecrById(id int64, col string, session *xorm.Session) (rs int64, err error) {
	m.UpdateAt = time.Now().UnixNano()
	if session != nil {
		rs, err = session.Where("id = ? ", id).Decr(col, 1).Update(m)
		if err != nil {
			return rs, errors.Wrap(err, "DecrById failed")
		}
		return rs, nil
	}
	rs, err = orm.SocialBotOrm.Cols(col).Where("id = ? ", id).Decr(col, 1).Update(m)
	if err != nil {
		return rs, errors.Wrap(err, "DecrById failed")
	}
	return rs, nil
}

func (m *Media) GetOneByUri(uri int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("uri=? and media_status = ? ", uri, common.MediaStatusPublished).Where("is_del=?", 0).Get(m)
	if err != nil {
		return rs, errors.Wrap(err, "GetOneByUri failed")
	}
	return rs, nil
}

func (m *Media) GetIdByUri(uri int64) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("uri=?", uri).Where("is_del=?", 0).Get(m)
	if err != nil {
		return rs, errors.Wrap(err, "GetIdByUri failed")
	}
	return rs, nil
}

func (ml *MediaList) GetListNewestByCid(lastId int64, category int) (err error) {
	where := "publish_at < ? and media_status = ?"
	orderBy := "publish_at desc"
	param := []interface{}{lastId, common.MediaStatusPublished}
	if lastId == 0 {
		where = "publish_at > ? and media_status = ?"
	}
	if category != 0 {
		where = where + " AND cid = ?"
		param = append(param, category)
	}
	err = orm.SocialBotOrm.Where(where, param...).Limit(common.DefaultPage).OrderBy(orderBy).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListNewestByCid failed")
	}
	return err
}

func (ml *MediaList) GetNewCreateByCid(category, limit int) (err error) {
	where := "media_status = ? and cid = ?"
	err = orm.SocialBotOrm.Where(where, common.MediaStatusPublished, category).Limit(limit).OrderBy("create_at desc").Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetNewCreateByCid failed")
	}
	return err
}

func (ml *MediaList) GetHotByCid(page int, category int) (err error) {
	where := "media_status = ?"
	orderBy := "like_num desc, view_num desc, publish_at desc"
	param := []interface{}{common.MediaStatusPublished}
	if category != 0 {
		where = where + " AND cid = ?"
		param = append(param, category)
	}
	err = orm.SocialBotOrm.Where(where, param...).Limit(common.DefaultPage, page*common.DefaultPage).OrderBy(orderBy).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetHotByCid failed")
	}
	return err
}

func (ml *MediaList) GetListByCidDESC(lastId int64, category int) (err error) {
	where := "publish_at < ? and media_status = ?"
	orderBy := "publish_at desc"
	param := []interface{}{lastId, common.MediaStatusPublished}
	if lastId == 0 {
		where = "publish_at > ? and media_status = ?"
	}
	if category != 0 {
		where = where + " AND cid = ?"
		param = append(param, category)
	}
	err = orm.SocialBotOrm.Where(where, param...).Limit(common.DefaultPage).OrderBy(orderBy).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListByCidDESC failed")
	}
	return err
}

func (ml *MediaList) GetListByCidASC(lastId int64, category int) (err error) {
	where := "id > ?  and media_status = ?"
	orderBy := "publish_at ASC"
	param := []interface{}{lastId, common.MediaStatusPublished}
	if category != 0 {
		where = where + " AND cid = ?"
		param = append(param, category)
	}
	err = orm.SocialBotOrm.Where(where, param...).Limit(common.DefaultPage).OrderBy(orderBy).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListByCidASC failed")
	}
	return err
}

func (ml *MediaList) GetListByMidList(midList []int64) (err error) {
	err = orm.SocialBotOrm.Omit("create_at", "update_at").In("id", midList).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("GetListByMidList failed, list %v", midList))
	}
	return err
}

func (ml *MediaList) GetListSortViews(page, limitNum int) (err error) {
	limit := common.DefaultPage
	param := []int{page * limit}
	if limitNum != 0 {
		limit = limitNum
		param = []int{}
	}
	err = orm.SocialBotOrm.Omit("create_at", "update_at").Desc("view_num").Limit(limit, param...).Where("is_del=?", 0).Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetListSortViews failed")
	}
	return err
}

func (ml *MediaList) GetRecommendList(lastId int64, limit int) (err error) {
	where := "recommend=? AND is_del=?"
	param := []interface{}{1, 0}
	if lastId != 0 {
		where = "recommend=? AND is_del=? AND id < ?"
		param = []interface{}{1, 0, lastId}
	}
	err = orm.SocialBotOrm.Where(where, param...).OrderBy("update_at DESC").Find(ml)
	if err != nil {
		return errors.Wrap(err, "GetRecommendList failed ")
	}
	return err
}
