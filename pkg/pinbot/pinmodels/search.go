package pinmodels

type SearchPinsResponse struct {
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Query     string   `json:"query"`
			Bookmarks []string `json:"bookmarks"`
			Scope     string   `json:"scope"`
		} `json:"options"`
	} `json:"resource"`
	ResourceResponse struct {
		ResponseError *ResponseError `json:"error"`
		Data          struct {
			Results []struct {
				GridDescription string `json:"grid_description"`
				TrackingParams  string `json:"tracking_params"`
				RichSummary     struct {
					Type     string `json:"type"`
					TypeName string `json:"type_name"`
				} `json:"rich_summary"`
				ID           string `json:"id"`
				IsPromoted   bool   `json:"is_promoted"`
				CommentCount int    `json:"comment_count"`
				Type         string `json:"type"`
				Method       string `json:"method"`
				GridTitle    string `json:"grid_title"`
				Description  string `json:"description"`
				IsPlayable   bool   `json:"is_playable"`
				Link         string `json:"link"`
				IsRepin      bool   `json:"is_repin"`
				IsUploaded   bool   `json:"is_uploaded"`
				IsNative     bool   `json:"is_native"`
				CreatedAt    string `json:"created_at"`
				RepinCount   int    `json:"repin_count"`
				IsVideo      bool   `json:"is_video"`
			} `json:"results"`
		} `json:"data"`
	} `json:"resource_response"`
}

type SearchMyPinsResponse struct {
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Query     string   `json:"query"`
			Bookmarks []string `json:"bookmarks"`
			Scope     string   `json:"scope"`
		} `json:"options"`
	} `json:"resource"`
	ResourceResponse struct {
		Data []struct {
			Domain          string      `json:"domain"`
			GridDescription string      `json:"grid_description"`
			Videos          interface{} `json:"videos"`
			TrackingParams  string      `json:"tracking_params"`
			StoryPinDataID  interface{} `json:"story_pin_data_id"`
			Images          struct {
				Seven36X struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"736x"`
				Four74X struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"474x"`
				One70X struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"170x"`
				Two36X struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"236x"`
				One36X136 struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"136x136"`
				Orig struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"orig"`
			} `json:"images"`
			ID                      string        `json:"id"`
			CarouselData            interface{}   `json:"carousel_data"`
			IsEligibleForWebCloseup bool          `json:"is_eligible_for_web_closeup"`
			IsPromoted              bool          `json:"is_promoted"`
			ShoppingFlags           []interface{} `json:"shopping_flags"`
			DescriptionHTML         string        `json:"description_html"`
			Privacy                 string        `json:"privacy"`
			BuyableProduct          interface{}   `json:"buyable_product"`
			Comments                struct {
				Bookmark interface{}   `json:"bookmark"`
				Data     []interface{} `json:"data"`
				URI      string        `json:"uri"`
			} `json:"comments"`
			Access       []string `json:"access"`
			CommentCount int      `json:"comment_count"`
			Board        struct {
				IsCollaborative  bool   `json:"is_collaborative"`
				Layout           string `json:"layout"`
				Name             string `json:"name"`
				Privacy          string `json:"privacy"`
				URL              string `json:"url"`
				CollaboratedByMe bool   `json:"collaborated_by_me"`
				Owner            struct {
					ID string `json:"id"`
				} `json:"owner"`
				FollowedByMe      bool   `json:"followed_by_me"`
				Type              string `json:"type"`
				ID                string `json:"id"`
				ImageThumbnailURL string `json:"image_thumbnail_url"`
			} `json:"board"`
			IsQuickPromotable       bool   `json:"is_quick_promotable"`
			Type                    string `json:"type"`
			Method                  string `json:"method"`
			IsWhitelistedForTriedIt bool   `json:"is_whitelisted_for_tried_it"`
			GridTitle               string `json:"grid_title"`
			ImageCrop               struct {
				MinY float64 `json:"min_y"`
				MaxY float64 `json:"max_y"`
			} `json:"image_crop"`
			ImageSignature                 string        `json:"image_signature"`
			Attribution                    interface{}   `json:"attribution"`
			Description                    string        `json:"description"`
			PriceValue                     float64       `json:"price_value"`
			PriceCurrency                  string        `json:"price_currency"`
			NativeCreator                  interface{}   `json:"native_creator"`
			IsPlayable                     bool          `json:"is_playable"`
			AdMatchReason                  int           `json:"ad_match_reason"`
			Link                           string        `json:"link"`
			HasRequiredAttributionProvider bool          `json:"has_required_attribution_provider"`
			ViewTags                       []interface{} `json:"view_tags"`
			IsRepin                        bool          `json:"is_repin"`
			DebugInfoHTML                  interface{}   `json:"debug_info_html"`
			IsUploaded                     bool          `json:"is_uploaded"`
			Pinner                         struct {
				Username               string `json:"username"`
				Type                   string `json:"type"`
				ExplicitlyFollowedByMe bool   `json:"explicitly_followed_by_me"`
				FullName               string `json:"full_name"`
				ImageSmallURL          string `json:"image_small_url"`
				BlockedByMe            bool   `json:"blocked_by_me"`
				ID                     string `json:"id"`
				ImageLargeURL          string `json:"image_large_url"`
			} `json:"pinner"`
			IsNative              bool          `json:"is_native"`
			CreatedAt             string        `json:"created_at"`
			RepinCount            int           `json:"repin_count"`
			Promoter              interface{}   `json:"promoter"`
			PromotedIsRemovable   bool          `json:"promoted_is_removable"`
			DominantColor         string        `json:"dominant_color"`
			DidIts                []interface{} `json:"did_its"`
			Title                 string        `json:"title"`
			Embed                 interface{}   `json:"embed"`
			AdditionalHideReasons []interface{} `json:"additional_hide_reasons"`
			IsVideo               bool          `json:"is_video"`
			IsDownstreamPromotion bool          `json:"is_downstream_promotion"`
		} `json:"data"`
		ResponseError *ResponseError `json:"error"`
	} `json:"resource_response"`
}

type SearchUsersResponse struct {
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Query     string   `json:"query"`
			Bookmarks []string `json:"bookmarks"`
			Scope     string   `json:"scope"`
		} `json:"options"`
	} `json:"resource"`
	ResourceResponse struct {
		Data struct {
			Typo interface{} `json:"typo"`
			Nag  struct {
			} `json:"nag"`
			ShouldAppendGlobalSearch  bool        `json:"should_append_global_search"`
			PersonalizationCandidates interface{} `json:"personalization_candidates"`
			Sensitivity               interface{} `json:"sensitivity"`
			Results                   []*UserItem `json:"results"`
		} `json:"data"`
		ResponseError *ResponseError `json:"error"`
	} `json:"resource_response"`
}
