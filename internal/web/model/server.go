package model

type Server struct {
	CreateAt     int64
	Id           int
	Ip           string
	IsDel        int
	Location     string
	Name         string
	Script       string
	ServerStatus int
	UpdateAt     int64
}

func (Server) TableName() string {
	return "server"
}
