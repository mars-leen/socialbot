package pincore

var (
	Header = map[string]string{
		"accept":               "application/json: text/javascript: */*: q=0.01",
		"accept-encoding":      "gzip: deflate: br",
		"Accept-Language":      "zh-CN:zh;q=0.9",
		"content-Type":         "application/x-www-form-urlencoded",
		"user-agent":           "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"x-app-version":        "61a7dfe",
		"x-pinterest-appstate": "active",
		"x-requested-with":     "XMLHttpRequest",
		"x-pinterest-mweb":     "1",
	}
)

const (
	CsrfTokenKey = "csrftoken"

	ReachSummaryProductType = "product"

	SearchScopePins   = "pins"
	SearchScopeMyPins = "my_pins"
	SearchScopeUsers  = "users"

	ApiAuthError                   = 3
	ApiErrorIpBanned               = 9
	ApiErrorAccountNotFound        = 79
	ApiErrorAccountInvalidPassword = 78
)

const (
	Host    = "www.pinterest.com"
	ApiHost = "https://www.pinterest.com/"

	// user
	ResourceLogin        = "resource/UserSessionResource/create/"     // login
	UserSettingsResource = "_ngjs/resource/UserSettingsResource/get/" // user setting
	UserResource         = "_ngjs/resource/UserResource/get/"         // user profile
	FollowUserResource   = "_ngjs/resource/UserFollowResource/create/"
	UnFollowUserResource = "_ngjs/resource/UserFollowResource/delete/"
	UserFansResource     = "_ngjs/resource/UserFollowersResource/get/" // user fan list

	// search
	BaseSearchResource = "resource/BaseSearchResource/get/" // search pins list and user list

	SearchResource = "resource/SearchResource/get/" // search my pins list

	// pin
	RepinResource = "_ngjs/resource/RepinResource/create/" // repin dest pin

	// board
	BoardsFeedResource = "_ngjs/resource/BoardsFeedResource/get/" // get user boards

)
