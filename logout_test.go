package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

func TestLogout(t *testing.T) {
	logoutResp, err := Logout(time.Second * 10)
	if err != nil {
		t.Fatal(err)
	}
	if !logoutResp.IsSuccessful() {
		t.Errorf("%+v", *logoutResp)
	}
}

func TestRawLogout(t *testing.T) {
	resp, err := logout(time.Second * 10)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // Ignore error.
	if err != nil {
		t.Fatal(err)
	}
	t.Log("len(body) =", len(body))
	t.Log(string(body))
	var result struct {
		Result  string
		Message string
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Result != "success" {
		t.Error(result.Result, "-", result.Message)
	}
}
