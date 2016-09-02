package web

import (
	"log"
	"net/http"
	"net/smtp"
)

//SendRequest Used To Send HTTP request to website
func SendRequest(url string) int {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}

func sendEmail(msgErr string, toEmail string, from string, password, string, mailServer string) {
	auth := smtp.PlainAuth("", from, password, mailServer)

	to := []string(toEmail)
	msg := []byte("to :recipient@example.net\r\n" +
		"Subject: Alert! non-200 response\r\n" +
		"\r\n" +
		msgErr)
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
