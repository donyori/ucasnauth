package main

import (
	"errors"
	"strings"
)

type TracePageError struct {
	pageName string
	msg      string
}

var ErrAlreadyLogin error = errors.New("ucasnauth: a user already login")
var ErrNotLogin error = errors.New("ucasnauth: no user login")

func NewTracePageError(pageName, msg string) error {
	pageName = strings.TrimSpace(pageName)
	msg = strings.TrimSpace(msg)
	msg = strings.TrimPrefix(msg, "ucasnauth:")
	msg = strings.TrimSpace(msg)
	return &TracePageError{pageName: pageName, msg: msg}
}

func (tpe *TracePageError) Error() string {
	return "ucasnauth: cannot trace " + tpe.pageName + " page - " + tpe.msg
}
