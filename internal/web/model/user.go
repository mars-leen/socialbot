package model

type User struct {
	Avatar     string
	CreateAt   int64
	Email      string
	Id         int64
	Identity   int
	Intro      string
	IsDel      int
	Nickname   string
	Password   string
	RegisterAt int64
	RegisterIp string
	UpdateAt   int64
	Uri        int64
	UserStatus int
}

func (User) TableName() string {
	return "user"
}
