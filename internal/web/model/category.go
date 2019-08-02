package model

type Category struct {
	Description string
	Id          int
	ShortName   string
	Sort        int
	Title       string
}

func (Category) TableName() string {
	return "category"
}
