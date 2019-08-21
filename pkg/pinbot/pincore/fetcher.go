package pincore

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ResponseCallback func(*http.Response) error

type Request struct {
	Url    string
	Method string
	Query  map[string]string
	Header map[string]string
}

type Fetcher struct{
	Client  *http.Client
	CsrfToken string
	ResponseCallback []ResponseCallback
}


func (f *Fetcher) BuildRequest(url, method, sourceUrl string, option map[string]interface{}) (*Request, error) {
	if f.CsrfToken == "" {
		return nil, errors.Errorf("url %s, csftoken is invalid", url)
	}

	// set request param
	data := map[string]interface{}{
		"options": option,
		"context": struct{}{},
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Json encode data: %v error", data))
	}
	param := map[string]string{
		"data":       string(dataJson),
		"source_url": sourceUrl,
	}

	// new request
	req := &Request{
		Url:    url,
		Method: method,
		Query:  param,
		Header: map[string]string{},
	}

	for key, value := range Header {
		req.Header[key]= value
	}
	req.Header["origin"] = strings.Trim(ApiHost, "/")
	req.Header["referer"] = ApiHost
	req.Header["x-csrftoken"] = f.CsrfToken
	/*req.Header.Set(":authority", HOST)
	req.Header.Set(":method", method)
	req.Header.Set(":path", u.Path)
	req.Header.Set(":scheme", u.Scheme)*/

	return req, nil
}

func (f *Fetcher) Get(urlStr string) (result []byte, err error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("invalid url(%s)", urlStr))
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("new requert url(%s) error ", urlStr))
	}
	resp, err := f.Client.Do(req)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("requert url(%s) error ", urlStr))
	}
	defer resp.Body.Close()

	// set token
	for _, value := range f.Client.Jar.Cookies(u) {
		if strings.Contains(value.Name, CsrfTokenKey) {
			f.CsrfToken = value.Value
		}
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("requert url(%s), read body error ", urlStr))
	}

	//check http status error
	err = CheckHttpsError(resp.StatusCode)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("requert url(%s), invalid http code, body %s", urlStr, string(result)))
	}
	return result, nil
}

func (f *Fetcher) Do(r *Request) (result []byte, err error) {
	u, err := url.Parse(r.Url)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("invalid url, param %#v", *r))
	}

	vs := url.Values{}
	bf := bytes.NewBuffer([]byte{})
	for k, v := range r.Query {
		vs.Add(k, v)
	}
	if r.Method == "POST" {
		bf.WriteString(vs.Encode())
	} else {
		for k, v := range u.Query() {
			vs.Add(k, strings.Join(v, " "))
		}
		u.RawQuery = vs.Encode()
	}

	req, err := http.NewRequest(r.Method, u.String(), bf)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("new request error, param %#v", *r))
	}

	for k, v := range r.Header{
		req.Header.Set(k, v)
	}

	resp, err := f.Client.Do(req)
	if err != nil {
		fmt.Println(resp.StatusCode)
		return result, errors.Wrap(err, fmt.Sprintf("do request error, resp %v, param %#v", *r, resp))
	}
	defer resp.Body.Close()

	// callback
	if len(f.ResponseCallback) > 0 {
		for _,fn := range f.ResponseCallback {
			err := fn(resp)
			if err != nil {
				return result, errors.Wrap(err, "ResponseCallback error")
			}
		}
	}

	// set token
	for _, value := range f.Client.Jar.Cookies(u) {
		if strings.Contains(value.Name, CsrfTokenKey) {
			f.CsrfToken = value.Value
		}
	}

	// gzip
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// read body
	result, err = ioutil.ReadAll(reader)
	if err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("ReadAll request body error, param %#v @ ", *r))
	}

	return result, nil
}


