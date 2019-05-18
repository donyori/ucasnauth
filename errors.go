package main

import (
	"errors"
	"strings"
)

type TraceLoginPageError struct {
	msg string
}

var ErrAlreadyLogin error = errors.New("ucasnauth: a user already login")
var ErrCannotGetMachineInfo error = errors.New("ucasnauth: cannot get machine information")

func NewTraceLoginPageError(msg string) error {
	msg = strings.TrimSpace(msg)
	msg = strings.TrimPrefix(msg, "ucasnauth:")
	msg = strings.TrimSpace(msg)
	return &TraceLoginPageError{msg: msg}
}

func (tlpe *TraceLoginPageError) Error() string {
	return "ucasnauth: cannot trace login page - " + tlpe.msg
}
