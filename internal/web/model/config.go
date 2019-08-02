package model

type Config struct {
	CreateAt int64
	Id       int
	IsDel    int
	Title    string
	UpdateAt int64
	Value    string
}

func (Config) TableName() string {
	return "config"
}
