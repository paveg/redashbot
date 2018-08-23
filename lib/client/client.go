package client

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

const (
	repositoryUrl   = "https://github.com/paveg/redashbot"
	redashUrlEnv    = "REDASH_URL"
	redashApikeyEnv = "REDASH_APIKEY"
)

type Options struct {
	Params map[string]string
	Header map[string]string
	Body   io.Reader
}

func defaultOptions() *Options {
	return &Options{
		Params: make(map[string]string),
		Header: map[string]string{
			"User-Agent": userAgent,
		},
		Body: nil,
	}
}

var (
	repository        = repositoryUrl
	userAgent         = fmt.Sprintf("RedashBot/0.1 (+%s; %s)", repository, runtime.Version())
	defaultPostHeader = map[string]string{
		"Content-Type": "application/json",
	}
	DefaultClient = NewDefaultClient()
)

type ClientData struct {
	*log.Logger
}

type DefaultClientData struct {
	ClientData
	apikey string
	url    *url.URL
}

func NewDefaultClient() *DefaultClientData {
	var err error
	var u *url.URL
	if ue := os.Getenv(redashUrlEnv); ue != "" {
		u, err = url.Parse(os.Getenv(redashUrlEnv))
		if err != nil {
			return nil
		}
		log.Printf("[Debug] set url: %s\n", u)
	} else {
		u = &url.URL{}
	}
	dcData := &DefaultClientData{
		apikey: os.Getenv(redashApikeyEnv),
		url:    u,
	}
	dcData.Logger = &log.Logger{}
	dcData.Logger.SetOutput(os.Stdout)
	dcData.Logger.SetFlags(log.Ldate | log.Ltime)
	return dcData
}
func (dc DefaultClientData) ApiKey() (apikey string, err error) {
	if len(dc.apikey) < 1 {
		dc.apikey = os.Getenv(redashApikeyEnv)
	}
	if len(dc.apikey) < 1 {
		return "", errors.New("Invalid apikey...")
	}
	dc.Logger.Printf("[DEBUG] apikey: [%s]\n", maskKey(dc.apikey))
	return dc.apikey, nil
}

func (dc DefaultClientData) DefaultOptions() *Options {
	return defaultOptions()
}

func (dc DefaultClientData) Url() (url *url.URL, err error) {
	if dc.url.String() == "" {
		dc.url, err = url.Parse(os.Getenv(redashUrlEnv))
		if err != nil {
			return nil, err
		}
	}
	return dc.url, err
}

func (dc DefaultClientData) HTTPClient() *http.Client {
	return http.DefaultClient
}

func maskKey(s string) string {
	var pre string
	if len(s) >= 4 {
		pre = s[0:4]
	} else {
		pre = "****"
	}
	return pre + "****"
}
