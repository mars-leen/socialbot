package model

type MediaTag struct {
	Cid            int
	CreateAt       int64
	Id             int64
	IsDel          int
	MediaPublishAt int64
	MediaStatus    int
	Mid            int64
	Tid            int
	UpdateAt       int64
}

func (MediaTag) TableName() string {
	return "media_tag"
}
