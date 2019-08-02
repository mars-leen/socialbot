package model

type MediaSource struct {
	CreateAt   int64
	Id         int64
	IsDel      int
	Mid        int64
	SourceType int
	UpdateAt   int64
	Url        string
}

func (MediaSource) TableName() string {
	return "media_source"
}
