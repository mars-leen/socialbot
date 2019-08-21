package pinbot

type Config struct {
	Proxy           string           `mapstructure:"proxy" json:"proxy"`
	UsernameOrEmail string           `mapstructure:"username_or_email" json:"username_or_email"`
	Password        string           `mapstructure:"password" json:"password"`
	FollowNumber    int              `mapstructure:"follow_number" json:"follow_number"`
	RepinNumber     int              `mapstructure:"repin_number" json:"repin_number"`
	Filter          FilterConfig     `mapstructure:"filter" json:"filter"`
	RandSleep       RandSleepConfig  `mapstructure:"rand_sleep" json:"rand_sleep"`
	SearchUser      SearchUserConfig `mapstructure:"search_user" json:"search_user"`
	SearchPin       SearchPinConfig  `mapstructure:"search_pin" json:"search_pin"`
}

type FilterConfig struct {
	FollowUserMinPin              int `mapstructure:"follow_user_min_pin" json:"follow_user_min_pin"`
	FollowUserMinFansSubFollowing int `mapstructure:"follow_user_min_fans_sub_following" json:"follow_user_min_fans_sub_following"`
}

type RandSleepConfig struct {
	Repin    []int `mapstructure:"repin" json:"repin"`
	Follow   []int `mapstructure:"follow" json:"follow"`
	Paginate []int `mapstructure:"paginate" json:"paginate"`
}

type SearchUserConfig struct {
	Paginate     int                      `mapstructure:"paginate" json:"paginate"`
	FansPaginate int                      `mapstructure:"fans_paginate" json:"fans_paginate"`
	Query        []*SearchUserQueryConfig `mapstructure:"query" json:"query"`
}

type SearchUserQueryConfig struct {
	Name                string `mapstructure:"name" json:"name"`
	SearchedUserMinFans int    `mapstructure:"searched_user_min_fans" json:"searched_user_min_fans"`
	SelectedUserNum     int    `mapstructure:"selected_user_num" json:"selected_user_num"`
}

type SearchPinConfig struct {
	Paginate int                     `mapstructure:"paginate" json:"paginate"`
	Query    []*SearchPinQueryConfig `json:"query" mapstructure:"query"`
}

type SearchPinQueryConfig struct {
	Name                string   `mapstructure:"name" json:"name"`
	RepinToBoard        string   `mapstructure:"repin_to_board" json:"repin_to_board"`
	SearchedPinMinRepin int      `mapstructure:"searched_pin_min_repin" json:"searched_pin_min_repin"`
	BookMarks           []string `mapstructure:"next_page" json:"next_page"`
}
