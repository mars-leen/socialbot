package model

type Crawler struct {
	CrawlerStatus int
	CreateAt      int64
	Id            int
	IsDel         int
	LastPage      string
	Name          string
	Script        string
}

func (Crawler) TableName() string {
	return "crawler"
}
