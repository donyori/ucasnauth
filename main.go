package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var cmdName string
	var resp Response
	var err error
	switch len(os.Args) {
	case 1:
		cmdName = "Authentication"
		resp, err = AuthWithLastInfo()
	case 2:
		cmd := strings.ToLower(os.Args[1])
		if cmd == "login" {
			cmdName = "Authentication"
			resp, err = AuthWithLastInfo()
		} else if cmd == "logout" {
			cmdName = "Logout"
			resp, err = DoLogout()
		} else {
			fmt.Println("Invalid arguments:", os.Args[1:])
			fmt.Println(UsageHint)
			return
		}
	case 3:
		if strings.ToLower(os.Args[1]) != "logout" {
			cmdName = "Authentication"
			resp, err = AuthWithGivenInfo(os.Args[1], os.Args[2])
		} else {
			cmdName = "Logout"
			resp, err = DoLogout()
		}
	case 4:
		cmd := strings.ToLower(os.Args[1])
		if cmd == "login" {
			cmdName = "Authentication"
			resp, err = AuthWithGivenInfo(os.Args[2], os.Args[3])
		} else if cmd == "logout" {
			cmdName = "Logout"
			resp, err = DoLogout()
		} else {
			fmt.Println("Invalid arguments:", os.Args[1:])
			fmt.Println(UsageHint)
			return
		}
	default:
		fmt.Println("Invalid arguments:", os.Args[1:])
		fmt.Println(UsageHint)
		return
	}
	if resp.IsSuccessful() {
		fmt.Println(cmdName, "succeeded.")
	} else if resp != nil {
		if msg := resp.GetMessage(); msg != "" {
			fmt.Println(cmdName, "failed. Message:", msg)
		} else {
			fmt.Println(cmdName, "failed.")
		}
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !resp.IsSuccessful() || err != nil {
		// Add a pause.
		fmt.Print("Press 'Enter' to exit...")
		// Pause at least one second.
		c := time.After(time.Second)
		doesContinue := true
		for doesContinue {
			fmt.Scanln()
			select {
			case <-c:
				doesContinue = false
			default:
			}
		}
	}
}
