package model

type Article struct {
	Content  string
	CreateAt int64
	Id       int64
	IsDel    int
	Mid      int64
	UpdateAt int64
}

func (Article) TableName() string {
	return "article"
}
