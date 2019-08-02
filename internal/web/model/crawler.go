package model

type Crawler struct {
	CrawlerStatus int
	CreateAt      int64
	Id            int
	IsDel         int
	LastPage      string
	Name          string
	Script        string
	UpdateAt      int64
}

func (Crawler) TableName() string {
	return "crawler"
}
