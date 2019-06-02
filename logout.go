package main

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type LogoutResp struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func Logout(timeout time.Duration) (logoutResp *LogoutResp, err error) {
	resp, err := logout(timeout)
	if err != nil {
		return
	}
	defer resp.Body.Close() // Ignore error.
	decoder := json.NewDecoder(resp.Body)
	result := new(LogoutResp)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	logoutResp = result
	return
}

func logout(timeout time.Duration) (*http.Response, error) {
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
		return nil, NewTracePageError("logout", "got "+traceResp.Status)
	}
	traceReqUrl := traceResp.Request.URL
	dir, file := path.Split(traceReqUrl.Path)
	if file == IndexPageBasename {
		return nil, ErrNotLogin
	} else if file != SuccessPageBasename {
		return nil, NewTracePageError("logout", "found "+file+
			" instead of "+SuccessPageBasename)
	}
	logoutUrl := &url.URL{
		Scheme:   traceReqUrl.Scheme,
		Host:     traceReqUrl.Host,
		Path:     path.Join(dir, InterFaceName),
		RawQuery: "method=logout",
	}
	logoutBody := "userIndex=" + traceReqUrl.Query().Get("userIndex")
	return client.Post(logoutUrl.String(), PostContentType,
		strings.NewReader(logoutBody))
}

func (lr *LogoutResp) IsSuccessful() bool {
	return lr != nil && strings.ToLower(lr.Result) == "success"
}

func (lr *LogoutResp) GetMessage() string {
	if lr == nil {
		return ""
	}
	return lr.Message
}
