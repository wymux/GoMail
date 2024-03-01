package main

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", "wymux@shoash.com", "password", "mail.shoash.com")
	to := []string{"wymux@shoash.com"}
		msg := []byte("To: recipient@example.net\r\n" +
		"Subject: Hello!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail("mail.shoash.com:25", auth, "wymux@shoash.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
