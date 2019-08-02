package model

type AccountForm struct {
	Account  string `form:"account" json:"account"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
}

type MediaAlbumForm struct {
	Type  int `form:"type" json:"type"`
	Medias string `form:"medias" json:"medias" binding:"required"`
	Desc string `form:"desc" json:"desc"`
	Tags string `form:"tags" json:"tags"`
	CategoryId int `form:"categoryId" json:"categoryId" binding:"required"`
	AlbumId int64 `form:"albumId" json:"albumId"`
}


type PublishConForm struct {
	Type  int `form:"type" json:"type"`
	Title string `form:"title" json:"title"`
	Desc string `form:"desc" json:"desc"`
	Cover string `form:"cover" json:"cover"`
	Medias string `form:"medias" json:"medias" binding:"required"`
	Tags string `form:"tags" json:"tags"`
	CategoryId int `form:"categoryId" json:"categoryId" binding:"required"`
	ContentId int64 `form:"contentId" json:"contentId"`
	Recommend bool `form:"recommend" json:"recommend"`
}


type CategoryForm struct {
	Id int `form:"id" json:"id" binding:"required"`
	ShortName string `form:"shortName" json:"shortName" binding:"required"`
	Title     string `form:"title" json:"title" binding:"required"`
	Description  string `form:"description" json:"description"  binding:"required"`
	Sort  int `form:"sort" json:"sort"  binding:"required"`
}

type TagForm struct {
	Id int `form:"id" json:"id" binding:"required"`
	Cid  int `form:"cid" json:"cid"  binding:"required"`
	ShortName string `form:"shortName" json:"shortName" binding:"required"`
	Title     string `form:"title" json:"title" binding:"required"`
	Description  string `form:"description" json:"description"  binding:"required"`
	BoardName string `form:"boardName" json:"boardName" binding:"required"`

}
