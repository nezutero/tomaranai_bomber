package main

import (
	"fmt"

	"github.com/kenjitheman/tomaranai_bomber/core"
)

func main() {
	fmt.Println(
		"[SUCCESS] tomaranai_bomber\n[!] usage:\n - type sms_bomb to start sms_bomb\n - type email_bomb to start email_bomb",
	)

	var arg string
	_, err := fmt.Scanf("%s", &arg)
	if err != nil {
		fmt.Println("[ERROR] incorrect input", err)
		return
	}

	switch arg {
	case "sms_bomb":
		core.SmsBomb()
	case "email_bomb":
		core.EmailBomb()
	default:
		fmt.Println("[ERROR] unknown option: ", arg)
	}
}
