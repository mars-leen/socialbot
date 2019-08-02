package model

type CCrawlerItem struct {
	Content     string
	Cover       string
	CreateAt    int64
	Crwid       int
	Description string
	Id          int64
	Title       string
}

func (CCrawlerItem) TableName() string {
	return "c_crawler_item"
}
