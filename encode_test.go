package main

import (
	"net/url"
	"testing"
)

func TestEncodeUriComponent(t *testing.T) {
	cases := []struct {
		Original string
		Encoded  string
	}{
		// Unchanged cases:
		{"", ""},
		{"true", "true"},
		{"-_.!~*'()", "-_.!~*'()"},

		// Cases from http://www.w3school.com.cn/jsref/jsref_encodeURIComponent.asp :
		{"http://www.w3school.com.cn", "http%3A%2F%2Fwww.w3school.com.cn"},
		{"http://www.w3school.com.cn/p 1/", "http%3A%2F%2Fwww.w3school.com.cn%2Fp%201%2F"},
		{",/?:@&=+$#", "%2C%2F%3F%3A%40%26%3D%2B%24%23"},
	}
	for _, c := range cases {
		s := EncodeUriComponent(c.Original)
		if s != c.Encoded {
			t.Error(s, "!=", c.Encoded)
		}
		// Compare with url.QueryEscape:
		q := url.QueryEscape(c.Original)
		if s != q {
			t.Log("encodeURIComponent:", s, "QueryEscape:", q)
		}
		// Compare with url.PathEscape:
		p := url.PathEscape(c.Original)
		if s != p {
			t.Log("encodeURIComponent:", s, "PathEscape:", p)
		}
	}
}
