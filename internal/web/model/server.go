package model

type Server struct {
	CreateAt     int64
	Id           int
	Ip           string
	Location     string
	Name         string
	Script       string
	ServerStatus int
}

func (Server) TableName() string {
	return "server"
}
