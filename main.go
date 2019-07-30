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
		switch cmd {
		case "login":
			cmdName = "Authentication"
			resp, err = AuthWithLastInfo()
		case "logout":
			cmdName = "Logout"
			resp, err = DoLogout()
		case "delete":
			cmdName = "Delete"
			resp, err = DoDelete()
		default:
			fmt.Println("Invalid arguments:", os.Args[1:])
			fmt.Println(UsageHint)
			return
		}
	case 3:
		arg1 := strings.ToLower(os.Args[1])
		if arg1 != "logout" && arg1 != "delete" {
			cmdName = "Authentication"
			resp, err = AuthWithGivenInfo(os.Args[1], os.Args[2])
		} else if arg1 == "logout" {
			cmdName = "Logout"
			resp, err = DoLogout()
		} else {
			cmdName = "Delete"
			resp, err = DoDelete()
		}
	case 4:
		cmd := strings.ToLower(os.Args[1])
		switch cmd {
		case "login":
			cmdName = "Authentication"
			resp, err = AuthWithGivenInfo(os.Args[2], os.Args[3])
		case "logout":
			cmdName = "Logout"
			resp, err = DoLogout()
		case "delete":
			cmdName = "Delete"
			resp, err = DoDelete()
		default:
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
	if resp.IsSuccessful() && err == nil {
		fmt.Println(cmdName, "succeeded.")
		// Sleep one second.
		time.Sleep(time.Second)
	} else {
		// Add a pause.
		fmt.Print("Press 'Enter' to exit...")
		// Pause at least one millisecond.
		c := time.After(time.Millisecond)
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
