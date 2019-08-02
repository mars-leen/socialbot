package model

type Pinbot struct {
	Config       string
	Id           int
	Name         string
	PinbotStatus int
	Sid          int
}

func (Pinbot) TableName() string {
	return "pinbot"
}
