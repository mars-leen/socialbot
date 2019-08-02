package model

type CrawlerItem struct {
	Content     string
	Cover       string
	CreateAt    int64
	Crwid       int
	Description string
	Id          int64
	IsDel       int
	Title       string
	UpdateAt    int64
}

func (CrawlerItem) TableName() string {
	return "crawler_item"
}
