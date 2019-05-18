package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestAuthenticate(t *testing.T) {
	// To test it, set a valid username and encryptedPassword below:
	username := ``
	encryptedPassword := ""
	if username == "" {
		return
	}
	testAuthenticate(t, username, encryptedPassword)
}

func TestRawAuthenticate(t *testing.T) {
	// To test it, set a valid username and encryptedPassword below:
	username := ``
	encryptedPassword := ""
	if username == "" {
		return
	}
	resp, err := authenticate(username, encryptedPassword, time.Second*10)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.Status)
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

func testAuthenticate(tb testing.TB, username, encryptedPassword string) {
	authResp, err := Authenticate(username, encryptedPassword, time.Second*10)
	if err != nil {
		tb.Fatal(err)
	}
	if !authResp.IsSuccessful() {
		tb.Errorf("%+v", *authResp)
	}
}
