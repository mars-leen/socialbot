package model

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
	Uid         int64
	UpdateAt    int64
	Uri         int64
	ViewNum     int
}

func (Media) TableName() string {
	return "media"
}
