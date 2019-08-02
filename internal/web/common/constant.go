package common

const (
	DefaultPage  = 15
	// account type
	AccountTypeEmail = "email"
	// jwt
	JwtAuthUserKey = "userId"
	SuperAdmin = 3

	// crawler
	ListPageNum = 10

	CrawlerContentNew = 0 // 新入库
	CrawlerContentPublished = 10 // 已发布

	CrawlerImgsLimit = 4 // 发布图片数量最大长度限制


	// media
	MediaTypeImg = 1 //"img"
	MediaTypeVideo = 2 //"video"
	MediaTypeArticle = 3 // "article"
	MediaPublished = 10
	MediaSourceFromLocal = 0 //local
	MediaSortNewest = 0
	MediaTagLimit = 100

	MediaUpAvatarDir = "avatar"
	MediaUpMediaDir = "media"
	MediaUpPublicDir = "public"

	// cache
	LikeMediaCache = "like_media_%d_%d"

	HostQiNiu = 1
	HostHuaBan =2

)