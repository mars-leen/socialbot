package pincore

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"socialbot/pkg/pinbot/pinmodels"
	"socialbot/pkg/utils"
)

func (p *Pinterest) AutoSetCsrfToken() error {
	// get csrfToken from cookie
	_, err := p.Fetcher.Get(ApiHost)
	if err != nil {
		return err
	}
	return err
}

func (p *Pinterest) LoginApi() (*pinmodels.LoginResponse, error) {
	// param
	api := ApiHost + ResourceLogin
	Method := "POST"
	sourceUrl := "/login/"
	option := map[string]interface{}{
		"username_or_email":    p.UsernameOrEmail,
		"password":             p.Password,
		"app_type_from_client": 6, // 6 mobile web 5 pc web
		"get_user":             true,
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send request
	resp, err := p.Fetcher.Do(req)
	if err != nil {
		return nil, err
	}

	// json decode
	respValue := &pinmodels.LoginResponse{}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp, respValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("json encode error %s", string(resp)))
	}

	// check error
	err = CheckRespErrors(respValue.ResourceResponse.ResponseError)
	if err != nil {
		return nil, err
	}

	// check banned
	if respValue.ResourceResponse.Data.User.IsWriteBanned {
		p.IsBanned = true
		return nil, errors.Wrap(err, "Congratulations ! your account has been banned")
	}

	p.IsLogin = true
	p.Username = respValue.ResourceResponse.Data.User.Username

	return respValue, nil
}

func (p *Pinterest) UserSettingApi() (*pinmodels.UserSettingsResponse, error) {
	// param
	api := ApiHost + UserSettingsResource
	Method := "GET"
	sourceUrl := "/settings/profile/"
	option := map[string]interface{}{}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.UserSettingsResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}

	// send request
	p.Username = respValue.ResourceResponse.Data.Username

	// check banned
	if respValue.ResourceResponse.Data.IsWriteBanned {
		p.IsBanned = true
		return nil, errors.Wrap(err, "Congratulations ! your account has been banned")
	}

	return respValue, nil
}

func (p *Pinterest) Send(req *Request, respValue interface{}) error {
	// send request
	resp, err := p.Fetcher.Do(req)
	if err != nil {
		return err
	}

	// json decode
	baseRsp := &pinmodels.BaseResponse{}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp, baseRsp)
	if err != nil {
		return errors.Wrap(err, string(resp))
	}

	// check error
	err = CheckRespErrors(baseRsp.ResourceResponse.ResponseError)
	if err == nil {
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp, respValue)
		if err != nil {
			return errors.Wrap(err, string(resp))
		}
		return  nil
	}

	// ##################################################
	fmt.Println("check and reLogin.....")
	err = p.CheckAndReLogin(err)
	if err != nil {
		return  err
	}
	fmt.Println("reLogin success then request .....")
	// ##################################################

	// send request
	resp, err = p.Fetcher.Do(req)
	if err != nil {
		return err
	}

	// json encode
	baseRsp = &pinmodels.BaseResponse{}
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp, baseRsp)
	if err != nil {
		return errors.Wrap(err, string(resp))
	}

	// check error
	err = CheckRespErrors(baseRsp.ResourceResponse.ResponseError)
	if err == nil {
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp, respValue)
		if err != nil {
			return errors.Wrap(err, string(resp))
		}
		return  nil
	}

	return err
}


//source_url: /mervatalashgar60/
//data: {"options":{"field_set_key":"profile","username":"mervatalashgar60","is_mobile_fork":true},"context":{}}
//_: 1564315865510
func (p *Pinterest) UserProfileApi(username string) (*pinmodels.UserProfileResponse, error) {
	// check
	if utils.TrimSpace(username) == "" {
		return nil, errors.New("username is empty")
	}

	// param
	api := ApiHost + UserResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/%s/", username)
	option := map[string]interface{}{
		"is_mobile_fork": "true",
		"field_set_key":  "profile",
		"username":       username,
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.UserProfileResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}

	return respValue, nil
}

func (p *Pinterest) FollowUserApi(userId, username string) (*pinmodels.FollowUserResponse, error) {
	// check
	if utils.TrimSpace(userId) == "" || utils.TrimSpace(username) == "" {
		return nil, errors.Errorf("username(%s) or userId(%s) is empty", username, userId)
	}

	// param
	api := ApiHost + FollowUserResource
	Method := "POST"
	sourceUrl := fmt.Sprintf("/%s/followers/", username)
	option := map[string]interface{}{
		"user_id": userId,
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.FollowUserResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}

	return respValue, nil
}

func (p *Pinterest) UnFollowUserApi(userId, username string) (*pinmodels.FollowUserResponse, error) {
	// check
	if utils.TrimSpace(userId) == "" || utils.TrimSpace(username) == "" {
		return nil, errors.Errorf("username(%s) or userId(%s) is empty", username, userId)
	}

	// param
	api := ApiHost + UnFollowUserResource
	Method := "POST"
	sourceUrl := fmt.Sprintf("/%s/followers/", username)
	option := map[string]interface{}{
		"user_id": userId,
	}

	// request
	req, err := p.Fetcher.BuildRequest(api, Method, sourceUrl, option)
	if err != nil {
		return nil, err
	}

	// send req
	respValue := &pinmodels.FollowUserResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}

	return respValue, nil
}

// explicitly_followed_by_me followed by me
func (p *Pinterest) UserFansApi(username string, bookmarks []string) (*pinmodels.UserFansResponse, error) {
	// check
	username = utils.TrimSpace(username)
	if username == ""{
		return nil, errors.New("username is empty!")
	}

	// param
	api := ApiHost+ UserFansResource
	Method := "GET"
	sourceUrl := fmt.Sprintf("/%s/followers/", username)
	option := map[string]interface{}{
		"username": username,
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
	respValue := &pinmodels.UserFansResponse{}
	err = p.Send(req, respValue)
	if  err != nil {
		return nil, err
	}
	return respValue, nil
}

func (p *Pinterest) CheckAndReLogin(err error) error {
	isAuthError := CheckIsAuthError(err)
	if !isAuthError {
		return err
	}
	fmt.Println("reLogin...............")
	_, err = p.LoginApi()
	return err
}

func (p *Pinterest) LoadCookieFromFile(u *url.URL) (bool, error){
	runPath, err := utils.RunPath()
	if err != nil {
		return false, errors.Wrap(err, "get RunPath error ")
	}
	fileName := utils.Md5(p.UsernameOrEmail + u.Path)
	fullPath := filepath.Join(runPath, "tmp", fileName)
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false, nil
	}

	re, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return  false, errors.Wrap(err, "read cookie file error")
	}

	var a []*http.Cookie
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(re, &a)
	if err != nil {
		return  false, errors.Wrap(err, fmt.Sprintf("unMarshal Json cookie error, str : %s", string(re)))
	}
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(u, a)
	p.Fetcher.Client.Jar = jar
	p.IsLogin = true
	return true, nil
}

func (p *Pinterest) SaveCookieToFile(u *url.URL, c []*http.Cookie) error {
	runPath, err := utils.RunPath()
	if err != nil {
		return errors.Wrap(err, "get RunPath error ")
	}
	tmpPath := filepath.Join(runPath, "tmp")
	fileName := utils.Md5(p.UsernameOrEmail + u.Path)
	fullPath := filepath.Join(tmpPath, fileName)
	_, err = os.Stat(tmpPath)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(tmpPath, os.ModePerm)
	}

	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "marshal cookie error ")
	}

	err = ioutil.WriteFile(fullPath, data, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Write cookie to File error ")
	}
	return  nil
}