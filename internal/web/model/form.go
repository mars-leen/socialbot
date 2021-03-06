package model

type AccountForm struct {
	Account  string `form:"account" json:"account"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
}

type CommissionProductForm struct {
	Id          int    `form:"id" json:"id"`
	Title       string `form:"title" json:"title" binding:"required"`
	Cover       string `form:"cover" json:"cover" `
	Medias      string `form:"medias" json:"medias" binding:"required"`
	Tags        string `form:"tags" json:"tags"`
	Cid         int    `form:"cid" json:"cid" binding:"required"`
	PromoteLink string `form:"promote_link" json:"promote_link" binding:"required"`
	DetailLink  string `form:"detail_link" json:"detail_link" binding:"required"`
}

type SocialProductForm struct {
	Id        int    `form:"id" json:"id"`
	Title     string `form:"title" json:"title" binding:"required"`
	Cover     string `form:"cover" json:"cover"`
	Medias    string `form:"medias" json:"medias" binding:"required"`
	Tags      string `form:"tags" json:"tags" binding:"required"`
	Cid       int    `form:"cid" json:"cid" binding:"required"`
	NeedFetch bool   `form:"needFetch" json:"needFetch"`
	Recommend bool   `form:"recommend" json:"recommend"`
	CrawlerItemId int64 `form:"crawlerItemId" json:"crawlerItemId"`
}

type EditMediaForm struct {
	Uri        int64    `form:"uri" json:"uri"`
	Title     string `form:"title" json:"title"`
	Tags      string `form:"tags" json:"tags"`
	Cid       int    `form:"cid" json:"cid"`
	Recommend bool   `form:"recommend" json:"recommend"`
}


type SubmitMediaForm struct {
	Uri      string  `json:"uri"`
	URL      string `json:"url"`
	FileType int    `json:"fileType"`
	FileName string `json:"fileName"`
	Source   int    `json:"source"`
}

type CategoryForm struct {
	Id          int    `form:"id" json:"id"`
	ShortName   string `form:"shortName" json:"shortName" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description"  binding:"required"`
	Sort        int    `form:"sort" json:"sort" `
}

type CopywriterForm struct {
	Id          int    `form:"id" json:"id"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description"  binding:"required"`
}

type TagForm struct {
	Id          int    `form:"id" json:"id"`
	Cid         int    `form:"cid" json:"cid"  binding:"required"`
	ShortName   string `form:"shortName" json:"shortName" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description"  binding:"required"`
	BoardName   string `form:"boardName" json:"boardName" binding:"required"`
}

type CrawlerForm struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	LastPage string `form:"last_page" json:"last_page"`
	Script   string `form:"script" json:"script"  binding:"required"`
}

type ConfigForm struct {
	Id    int    `form:"id" json:"id"`
	Key   string `form:"key" json:"key"`
	Title string `form:"title" json:"title" binding:"required"`
	Value string `form:"value" json:"value"  binding:"required"`
	Sort  int    `form:"sort" json:"sort" `
}
