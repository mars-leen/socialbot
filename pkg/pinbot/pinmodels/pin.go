package pinmodels

type RePinResponse struct {
	ResourceResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Code    int    `json:"code"`
		HTTPStatus int `json:"http_status"`
		ResponseError *ResponseError `json:"error"`
		Data struct {
			Domain string `json:"domain"`
			IsQuickPromotable bool `json:"is_quick_promotable"`
			Description string `json:"description"`
			Pinner struct {
				Username string `json:"username"`
				ExplicitlyFollowedByMe bool `json:"explicitly_followed_by_me"`
				FullName string `json:"full_name"`
				FollowerCount int `json:"follower_count"`
				ImageSmallURL string `json:"image_small_url"`
				ID string `json:"id"`
			} `json:"pinner"`
			TrackingParams string `json:"tracking_params"`
			CreatedAt string `json:"created_at"`
			IsQuickPromotableByPinner bool `json:"is_quick_promotable_by_pinner"`
			ID string `json:"id"`
			Images struct {
				Four74X struct {
					URL string `json:"url"`
					Width int `json:"width"`
					Height int `json:"height"`
				} `json:"474x"`
				Two36X struct {
					URL string `json:"url"`
					Width int `json:"width"`
					Height int `json:"height"`
				} `json:"236x"`
			} `json:"images"`
			Link string `json:"link"`
			Board struct {
				Category string `json:"category"`
				IsCollaborative bool `json:"is_collaborative"`
				Name string `json:"name"`
				Privacy string `json:"privacy"`
				URL string `json:"url"`
				ImageCoverURL string `json:"image_cover_url"`
				Images struct {
					One70X []struct {
						URL string `json:"url"`
						Width int `json:"width"`
						DominantColor string `json:"dominant_color"`
						Height int `json:"height"`
					} `json:"170x"`
				} `json:"images"`
				Type string `json:"type"`
				ID string `json:"id"`
			} `json:"board"`
			Section interface{} `json:"section"`
			LinkDomain struct {
				OfficialUser struct {
					Username string `json:"username"`
					ExplicitlyFollowedByMe bool `json:"explicitly_followed_by_me"`
					FullName string `json:"full_name"`
					FollowerCount int `json:"follower_count"`
					ImageSmallURL string `json:"image_small_url"`
					ID string `json:"id"`
				} `json:"official_user"`
			} `json:"link_domain"`
			RichSummary struct {
				ID string `json:"id"`
			} `json:"rich_summary"`
			Type string `json:"type"`
			IsRepin bool `json:"is_repin"`
		} `json:"data"`
	} `json:"resource_response"`
}
