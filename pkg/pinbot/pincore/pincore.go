package pincore

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"socialbot/pkg/utils"
	"time"
)

type Pinterest struct {
	UsernameOrEmail string
	Password        string
	Username        string
	CsrfToken       string
	Fetcher         *Fetcher
	IsLogin         bool
	IsBanned         bool
}

func New(fns ...func(p *Pinterest) error) (*Pinterest, error) {
	jar, _ := cookiejar.New(nil)
	pin := &Pinterest{
		Fetcher: &Fetcher{
			Client:&http.Client{
				Timeout: 120 * time.Second,
				Jar:     jar,
			},
		},
	}
	for _, fn := range fns {
		err := fn(pin)
		if err != nil {
			return nil, err
		}
	}
	return pin, nil
}

// set username and password
func SetUsernamePassword(usernameOrEmail, password string) func(p *Pinterest) error {
	return func(p *Pinterest) error {
		p.UsernameOrEmail = usernameOrEmail
		p.Password = password
		return nil
	}
}

// set (n) second timeout
func SetRequestTimeout(n time.Duration) func(p *Pinterest) error {
	return func(p *Pinterest) error {
		p.Fetcher.Client.Timeout = n * time.Second
		return nil
	}
}

// set proxy, eg: http://127.0.0.1:8888
// not recommend to switch proxy, one user one proxy
func SetProxy(urlStr string) func(p *Pinterest) error {
	return func(p *Pinterest) error {
		if utils.TrimSpace(urlStr) == ""{
			return nil
		}
		u, err := url.Parse(urlStr)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("parse proxy url %s err", urlStr))
		}
		p.Fetcher.Client.Transport = &http.Transport{
			Proxy: http.ProxyURL(u),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		return nil
	}
}

