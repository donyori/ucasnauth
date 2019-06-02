package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type AuthResp struct {
	UserIndex         string `json:"userIndex"`
	Result            string `json:"result"`
	Message           string `json:"message"`
	ForWordUrl        string `json:"forwordurl"`
	KeepaliveInterval int    `json:"keepaliveInterval"`
	ValidCodeUrl      string `json:"validCodeUrl"`
}

func Authenticate(username, encryptedPassword string, timeout time.Duration) (
	authResp *AuthResp, err error) {
	resp, err := authenticate(username, encryptedPassword, timeout)
	if err != nil {
		return
	}
	defer resp.Body.Close() // Ignore error.
	decoder := json.NewDecoder(resp.Body)
	result := new(AuthResp)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	authResp = result
	return
}

func authenticate(username, encryptedPassword string, timeout time.Duration) (
	*http.Response, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}
	traceResp, err := client.Get(HostUrl)
	if err != nil {
		return nil, err
	}
	defer traceResp.Body.Close()
	if traceResp.StatusCode != http.StatusOK {
		return nil, NewTracePageError("login", "got "+traceResp.Status)
	}
	traceReqUrl := traceResp.Request.URL
	dir, file := path.Split(traceReqUrl.Path)
	if file == SuccessPageBasename {
		return nil, ErrAlreadyLogin
	} else if file != IndexPageBasename {
		return nil, NewTracePageError("login", "found "+file+
			" instead of "+IndexPageBasename)
	}
	authUrl := &url.URL{
		Scheme:   traceReqUrl.Scheme,
		Host:     traceReqUrl.Host,
		Path:     path.Join(dir, InterFaceName),
		RawQuery: "method=login",
	}
	authBody := fmt.Sprintf(AuthBodyFormat, EncodeUriComponentTwice(username),
		encryptedPassword, EncodeUriComponentTwice(traceReqUrl.RawQuery))
	return client.Post(authUrl.String(), PostContentType,
		strings.NewReader(authBody))
}

func (ar *AuthResp) IsSuccessful() bool {
	return ar != nil && strings.ToLower(ar.Result) == "success"
}

func (ar *AuthResp) GetMessage() string {
	if ar == nil {
		return ""
	}
	return ar.Message
}
