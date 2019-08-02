package model

type Pinbot struct {
	Config       string
	CreateAt     int64
	Id           int
	IsDel        int
	Name         string
	PinbotStatus int
	Sid          int
	UpdateAt     int64
}

func (Pinbot) TableName() string {
	return "pinbot"
}
