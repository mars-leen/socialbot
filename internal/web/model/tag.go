package model

type Tag struct {
	BoardName   string
	Cid         int
	Description string
	Id          int
	ShortName   string
	Title       string
}

func (Tag) TableName() string {
	return "tag"
}
