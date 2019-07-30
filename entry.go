package main

import (
	"encoding/json"
	"strings"
	"time"
)

type DeleteResp struct {
	isSuccessful bool
	msg          string
}

type usernameAndPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const defaultTimeout time.Duration = time.Second * 10

func AuthWithGivenInfo(username, password string) (*AuthResp, error) {
	username = strings.TrimSpace(username)
	pubKey := NewRsaPublicKey(RsaPubKeyExp, RsaPubKeyModHex, 0)
	encrypted := RsaEncrypt(pubKey, EncodeUriComponentTwice(password), 16)
	authResp, err := Authenticate(username, encrypted, defaultTimeout)
	if err != nil {
		return nil, err
	}
	if authResp.IsSuccessful() {
		d := &usernameAndPassword{
			Username: username,
			Password: encrypted,
		}
		var data []byte
		data, err = json.Marshal(d)
		if err == nil {
			err = Save(data)
		}
	}
	return authResp, err
}

func AuthWithLastInfo() (*AuthResp, error) {
	data, err := Load()
	if err != nil {
		return nil, err
	}
	d := new(usernameAndPassword)
	err = json.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}
	return Authenticate(d.Username, d.Password, defaultTimeout)
}

func DoLogout() (*LogoutResp, error) {
	return Logout(defaultTimeout)
}

func DoDelete() (*DeleteResp, error) {
	err := Delete()
	resp := new(DeleteResp)
	if err == nil {
		resp.isSuccessful = true
	} else {
		resp.msg = err.Error()
	}
	return resp, nil
}

func (dr *DeleteResp) IsSuccessful() bool {
	return dr != nil && dr.isSuccessful
}

func (dr *DeleteResp) GetMessage() string {
	if dr == nil {
		return ""
	}
	return dr.msg
}
