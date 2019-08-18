package common

const (
	DefaultPage = 15
	GalleryPage = 30
	SortNew    = 0
	sortHot   = 1
	// account type
	AccountTypeEmail = "email"
	// jwt
	JwtAuthUserKey = "userId"
	SuperAdmin     = 3

	// storage
	StorageMediaDir      = "media"
	StorageAvatarDir      = "avatar"
	StoragePublicDir      = "public"
	DefaultStorageSource = 0
	SourceTypeImage      = 0
	SourceTypeVideo      = 1

	// crawler
	CrawlerListPageNum   = 2
	CrawlerItemNew       = 0  // 新入库
	CrawlerItemPublished = 10 // 已发布

	// media
	MediaStatusPublished      = 10
	MediaTypePromotionProduct = 4
	MediaTypeArticle          = 3
	MediaTypeSocial           = 2
)
