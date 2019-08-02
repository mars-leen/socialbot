package model

type Article struct {
	Content string
	Id      int64
	IsDel   int
	Mid     int64
}

func (Article) TableName() string {
	return "article"
}
