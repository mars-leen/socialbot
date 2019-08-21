package pinmodels

type UserItem struct {
	Username               string   `json:"username"`
	DomainVerified         bool     `json:"domain_verified"`
	BlockedByMe            bool     `json:"blocked_by_me"`
	Type                   string   `json:"type"`
	ExplicitlyFollowedByMe bool     `json:"explicitly_followed_by_me"`
	FullName               string   `json:"full_name"`
	FollowerCount          int      `json:"follower_count"`
	VerifiedIdentity       struct {} `json:"verified_identity"`
	PinCount        int    `json:"pin_count"`
	ID              string `json:"id"`
	BoardCount      int    `json:"board_count"`
}

type LoginResponse struct {
	ResourceResponse struct {
		ResponseError *ResponseError `json:"error"`
		Data struct {
			TokenType string `json:"token_type"`
			User      struct {
				IsEmployee                          bool          `json:"is_employee"`
				Username                            string        `json:"username"`
				PhoneNumberEnd                      string        `json:"phone_number_end"`
				LastName                            string        `json:"last_name"`
				DomainVerified                      bool          `json:"domain_verified"`
				Locale                              string        `json:"locale"`
				PushPackageUserID                   string        `json:"push_package_user_id"`
				HasPassword                         bool          `json:"has_password"`
				FullName                            string        `json:"full_name"`
				ID                                  string        `json:"id"`
				IsWriteBanned                       bool          `json:"is_write_banned"`
				FirstName                           string        `json:"first_name"`
				DomainURL                           interface{}   `json:"domain_url"`
				ProfileDiscoveredPublic             interface{}   `json:"profile_discovered_public"`
				UnverifiedPhoneNumber               interface{}   `json:"unverified_phone_number"`
				Type                                string        `json:"type"`
				Email                               string        `json:"email"`
				ImageLargeURL                       string        `json:"image_large_url"`
				PhoneNumber                         interface{}   `json:"phone_number"`
				CanEnableMfa                        bool          `json:"can_enable_mfa"`
				PhoneCountry                        interface{}   `json:"phone_country"`
				VerifiedIdentity                    struct {} `json:"verified_identity"`
				UnverifiedPhoneCountry             interface{}   `json:"unverified_phone_country"`
				HasMfaEnabled                      bool          `json:"has_mfa_enabled"`
				TwitterPublishEnabled              bool          `json:"twitter_publish_enabled"`
				ResurrectionInfo                   interface{}   `json:"resurrection_info"`
				IsPartner                          bool          `json:"is_partner"`
				IsHighRisk                         bool          `json:"is_high_risk"`
			} `json:"user"`
		} `json:"data"`
	} `json:"resource_response"`
}

type UserSettingsResponse struct {
	ResourceResponse struct {
		ResponseError *ResponseError `json:"error"`
		Data    struct {
			ImpressumURL                       interface{}   `json:"impressum_url"`
			LastName                           string        `json:"last_name"`
			EmailSettings                      string        `json:"email_settings"`
			Locale                             string        `json:"locale"`
			HasPassword                        bool          `json:"has_password"`
			ThirdPartyMarketingTrackingEnabled bool          `json:"third_party_marketing_tracking_enabled"`
			CustomGender                       interface{}   `json:"custom_gender"`
			Gender                             string        `json:"gender"`
			NewsSettings                       string        `json:"news_settings"`
			ID                                 string        `json:"id"`
			IsWriteBanned                      bool          `json:"is_write_banned"`
			FirstName                          string        `json:"first_name"`
			PushSettings                       string        `json:"push_settings"`
			PersonalizeFromOffsiteBrowsing     bool          `json:"personalize_from_offsite_browsing"`
			FacebookTimelineEnabled            bool          `json:"facebook_timeline_enabled"`
			EmailChangingTo                    interface{}   `json:"email_changing_to"`
			PersonalizeNuxFromOffsiteBrowsing  bool          `json:"personalize_nux_from_offsite_browsing"`
			HasConfirmedEmail                  bool          `json:"has_confirmed_email"`
			IsPartner                          bool          `json:"is_partner"`
			Type                               string        `json:"type"`
			Email                              string        `json:"email"`
			WebsiteURL                         string        `json:"website_url"`
			Location                           string        `json:"location"`
			Username                           string        `json:"username"`
			About                              string        `json:"about"`
			ProfileImageURL                    string        `json:"profile_image_url"`
			EmailBounced                       bool          `json:"email_bounced"`
			AdsCustomizeFromConversion         bool          `json:"ads_customize_from_conversion"`
			AdditionalWebsiteUrls              []interface{} `json:"additional_website_urls"`
			FacebookPublishStreamEnabled       bool          `json:"facebook_publish_stream_enabled"`
			IsHighRisk                         bool          `json:"is_high_risk"`
			ShowImpressum                      bool          `json:"show_impressum"`
			Age                                interface{}   `json:"age"`
			ExcludeFromSearch                  bool          `json:"exclude_from_search"`
			Birthdate                          interface{}   `json:"birthdate"`
			PfyPreference                      bool          `json:"pfy_preference"`
			EmailBizSettings                   string        `json:"email_biz_settings"`
			Country                            string        `json:"country"`
			HideFromNews                       bool          `json:"hide_from_news"`
		} `json:"data"`
	} `json:"resource_response"`
}

type UserProfileResponse struct {
	ResourceResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Code    int    `json:"code"`
		HTTPStatus int `json:"http_status"`
		ResponseError *ResponseError `json:"error"`
		Data    struct {
			ShowCreatorProfile          bool        `json:"show_creator_profile"`
			ImpressumURL                interface{} `json:"impressum_url"`
			LastName                    string      `json:"last_name"`
			DomainVerified              bool        `json:"domain_verified"`
			FollowingCount              int         `json:"following_count"`
			IsTastemaker                bool        `json:"is_tastemaker"`
			ProfileReach                int         `json:"profile_reach"`
			ImageMediumURL              string      `json:"image_medium_url"`
			StoryPinCount               int         `json:"story_pin_count"`
			FullName                    string      `json:"full_name"`
			ID                          string      `json:"id"`
			StorefrontManagementEnabled bool        `json:"storefront_management_enabled"`
			FirstName                   string      `json:"first_name"`
			DomainURL                   interface{} `json:"domain_url"`
			HasShoppingShowcase         bool        `json:"has_shopping_showcase"`
			ExplicitlyFollowedByMe      bool        `json:"explicitly_followed_by_me"`
			HasCatalog                  bool        `json:"has_catalog"`
			Location                    string      `json:"location"`
			Indexed                     bool        `json:"indexed"`
			ProfileDiscoveredPublic     interface{} `json:"profile_discovered_public"`
			Type                        string      `json:"type"`
			WebsiteURL                  interface{} `json:"website_url"`
			BoardCount                  int         `json:"board_count"`
			Username                    string      `json:"username"`
			VerifiedIdentity            struct {
			} `json:"verified_identity"`
			VideoPinCount              int         `json:"video_pin_count"`
			HasShowcase                bool        `json:"has_showcase"`
			LastPinSaveTime            string      `json:"last_pin_save_time"`
			FollowerCount              int         `json:"follower_count"`
			IsPartner                  bool        `json:"is_partner"`
			PinsDoneCount              int         `json:"pins_done_count"`
			HasCustomBoardSortingOrder bool        `json:"has_custom_board_sorting_order"`
			PinCount                   int         `json:"pin_count"`
			NativePinCount             int         `json:"native_pin_count"`
			BlockedByMe bool `json:"blocked_by_me"`
			HasBoard bool `json:"has_board"`
		} `json:"data"`
	} `json:"resource_response"`
}

type FollowUserResponse struct {
	ResourceResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Code    int    `json:"code"`
		HTTPStatus int `json:"http_status"`
		ResponseError *ResponseError `json:"error"`
		Data    struct {
			Username       string `json:"username"`
			FirstName      string `json:"first_name"`
			LastName       string `json:"last_name"`
			Gender         string `json:"gender"`
			ImageMediumURL string `json:"image_medium_url"`
			ImageXlargeURL string `json:"image_xlarge_url"`
			FullName       string `json:"full_name"`
			ImageSmallURL  string `json:"image_small_url"`
			Type           string `json:"type"`
			ID             string `json:"id"`
			ImageLargeURL  string `json:"image_large_url"`
		} `json:"data"`
	} `json:"resource_response"`
}

type UserFansResponse struct {
	ResourceResponse struct {
		Data     []*UserItem `json:"data"`
		ResponseError *ResponseError `json:"error"`
	} `json:"resource_response"`
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Bookmarks []string `json:"bookmarks"`
			Username  string   `json:"username"`
		} `json:"options"`
	} `json:"resource"`
}

