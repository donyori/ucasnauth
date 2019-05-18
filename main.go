package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var authResp *AuthResp
	var err error
	if len(os.Args) == 1 {
		authResp, err = AuthWithLastInfo()
	} else if len(os.Args) == 3 {
		authResp, err = AuthWithGivenInfo(os.Args[1], os.Args[2])
		if strings.ContainsRune(os.Args[1], '\\') {
			fmt.Println(`Error: username cannot contain "\"! ... Just a joke, please ignore this. :)`)
		}
	} else {
		fmt.Println("Invalid arguments:", os.Args[1:])
		fmt.Println(UsageHint)
		return
	}
	if authResp.IsSuccessful() {
		fmt.Println("Authentication succeeded.")
	} else if authResp != nil {
		if authResp.Message != "" {
			fmt.Println("Authentication failed. Message:", authResp.Message)
		} else {
			fmt.Println("Authentication failed.")
		}
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
}
