package pincore

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/pkg/pinbot/pinmodels"
	"socialbot/pkg/utils"
)

func (p *Pinterest) UserBoardsApi(username string) (*pinmodels.UserBoardsResponse, error) {
	// check
	if utils.TrimSpace(username) == "" {
		return nil, errors.New("username is empty")
	}

	// param
	api := ApiHost+BoardsFeedResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/%s/", username)
	option := map[string]interface{}{
		"sort": "last_pinned_to",
		"field_set_key": "profile_grid_item",
		"username": username,
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.UserBoardsResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}

