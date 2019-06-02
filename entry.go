package main

import (
	"encoding/json"
	"strings"
	"time"
)

type usernameAndPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthWithGivenInfo(username, password string) (*AuthResp, error) {
	username = strings.TrimSpace(username)
	pubKey := NewRsaPublicKey(RsaPubKeyExp, RsaPubKeyModHex, 0)
	encrypted := RsaEncrypt(pubKey, EncodeUriComponentTwice(password), 16)
	return authAndUpdateData(username, encrypted)
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
	return authAndUpdateData(d.Username, d.Password)
}

func DoLogout() (*LogoutResp, error) {
	return Logout(time.Second * 15)
}

func authAndUpdateData(username, encryptedPassword string) (*AuthResp, error) {
	authResp, err := Authenticate(username, encryptedPassword, time.Second*15)
	if err != nil {
		return nil, err
	}
	if authResp.IsSuccessful() {
		d := &usernameAndPassword{
			Username: username,
			Password: encryptedPassword,
		}
		var data []byte
		data, err = json.Marshal(d)
		if err == nil {
			err = Save(data)
		}
	}
	return authResp, err
}
