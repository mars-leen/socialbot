package pinmodels
type UserBoardsResponse struct {
	ResourceResponse struct {
		ResponseError *ResponseError `json:"error"`
		Data    []struct {
			CollaboratingUsers          []interface{} `json:"collaborating_users"`
			ImageCoverURL               string        `json:"image_cover_url"`
			CollaboratorRequestsEnabled bool          `json:"collaborator_requests_enabled"`
			ArchivedByMeAt              interface{}   `json:"archived_by_me_at"`
			Owner                       struct {
				Username               string `json:"username"`
				DomainVerified         bool   `json:"domain_verified"`
				IsDefaultImage         bool   `json:"is_default_image"`
				ImageMediumURL         string `json:"image_medium_url"`
				ExplicitlyFollowedByMe bool   `json:"explicitly_followed_by_me"`
				FullName               string `json:"full_name"`
				VerifiedIdentity       struct {
				} `json:"verified_identity"`
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"owner"`
			Images struct {
				One70X []struct {
					URL           string `json:"url"`
					Width         int    `json:"width"`
					DominantColor string `json:"dominant_color"`
					Height        int    `json:"height"`
				} `json:"170x"`
			} `json:"images"`
			ID                           string   `json:"id"`
			ImageCoverHdURL              string   `json:"image_cover_hd_url"`
			Privacy                      string   `json:"privacy"`
			HasCustomCover               bool     `json:"has_custom_cover"`
			Access                       []string `json:"access"`
			FollowerCount                int      `json:"follower_count"`
			FollowedByMe                 bool     `json:"followed_by_me"`
			Type                         string   `json:"type"`
			SectionCount                 int      `json:"section_count"`
			IsCollaborative              bool     `json:"is_collaborative"`
			Description                  string   `json:"description"`
			BoardOrderModifiedAt         string   `json:"board_order_modified_at"`
			CollaboratorCount            int      `json:"collaborator_count"`
			AllowHomefeedRecommendations bool     `json:"allow_homefeed_recommendations"`
			CollaboratedByMe             bool     `json:"collaborated_by_me"`
			PinCount                     int      `json:"pin_count"`
			CoverPin                     struct {
				PinID          string `json:"pin_id"`
				ImageSignature string `json:"image_signature"`
				Timestamp      int    `json:"timestamp"`
			} `json:"cover_pin"`
			Name        string `json:"name"`
			CoverImages struct {
				Two22X struct {
					URL    string      `json:"url"`
					Width  int         `json:"width"`
					Height interface{} `json:"height"`
				} `json:"222x"`
			} `json:"cover_images"`
			URL                             string `json:"url"`
			CreatedAt                       string `json:"created_at"`
			ShouldShowMoreIdeas             bool   `json:"should_show_more_ideas"`
			ViewerCollaboratorJoinRequested bool   `json:"viewer_collaborator_join_requested"`
		} `json:"data"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Code    int    `json:"code"`
		HTTPStatus int `json:"http_status"`
	} `json:"resource_response"`
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Bookmarks   []string `json:"bookmarks"`
			Sort        string   `json:"sort"`
			FieldSetKey string   `json:"field_set_key"`
			Username    string   `json:"username"`
		} `json:"options"`
	} `json:"resource"`
}
