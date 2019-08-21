package pincore

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"socialbot/pkg/pinbot/pinmodels"
	"socialbot/pkg/utils"
)

func (p *Pinterest) SearchPinsApi(query string, bookmarks []string) (*pinmodels.SearchPinsResponse, error) {
	// check
	query = utils.TrimSpace(query)
	if query == ""{
		return nil, errors.New("search query is empty!")
	}

	// param
	api := ApiHost+BaseSearchResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/search/pins/?rs=ac&len=2&q=%s", url.QueryEscape(query))
	option := map[string]interface{}{
		"query": query,
		"scope": SearchScopePins,
	}
	if len(bookmarks) != 0 {
		Method = "POST"
		option["bookmarks"] = bookmarks
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.SearchPinsResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}

func (p *Pinterest) SearchUsersApi(query string, bookmarks []string) (*pinmodels.SearchUsersResponse, error) {
	// check
	query = utils.TrimSpace(query)
	if query == ""{
		return nil, errors.New("search query is empty!")
	}

	// param
	api := ApiHost+BaseSearchResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/search/users/?rs=ac&len=2&q=%s", url.QueryEscape(query))
	option := map[string]interface{}{
		"query": query,
		"scope": SearchScopeUsers,
	}
	if len(bookmarks) != 0 {
		option["bookmarks"] = bookmarks
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.SearchUsersResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}

func (p *Pinterest) SearchMyPinsApi(query, bookmarks string) (*pinmodels.SearchMyPinsResponse, error) {
	// check
	query = utils.TrimSpace(query)
	bookmarks = utils.TrimSpace(bookmarks)
	if query == ""{
		return nil, errors.New("search query is empty!")
	}

	// param
	api := ApiHost+ SearchResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/search/my_content/?rs=ac&len=2&q=%s", url.QueryEscape(query))
	option := map[string]interface{}{
		"query": query,
		"scope": SearchScopeMyPins,
	}
	if bookmarks != "" {
		option["bookmarks"] = bookmarks
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.SearchMyPinsResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}