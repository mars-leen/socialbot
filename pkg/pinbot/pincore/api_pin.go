package pincore

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/pkg/pinbot/pinmodels"
)

func (p *Pinterest) RePinApi(pinId, borderId, description, link, title, trackingParam string) (*pinmodels.RePinResponse, error)  {
	// check
	if pinId == "" || borderId == ""{
		return nil, errors.Errorf("pinId(%s) Or borderId(%s)  is empty", pinId, borderId)
	}

	// param
	api := ApiHost+RepinResource
	Method := "POST"
	sourceUrl := fmt.Sprintf("/pin/%s/repin/", pinId)
	option := map[string]interface{}{
		"pin_id": pinId,
		"board_id": borderId,
		"description": description,
		"clientTrackingParams": pinId,
		"link": borderId,
		"title": description,
		"is_removable": struct {}{},
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.RePinResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}