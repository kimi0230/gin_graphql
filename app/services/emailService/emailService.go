package emailService

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

var smtp string
var user string
var password string
var tlsPort int

type EmailStruct struct {
	From    string
	To      string
	Cc      string
	Subject string
	Body    string
	Attach  string
}

func init() {
	smtp = os.Getenv("GMail_SMTP")
	user = os.Getenv("GMail_User")
	password = os.Getenv("GMail_Password")
	tlsPort, _ = strconv.Atoi(os.Getenv("GMail_TLS_Port"))
}

func Send(es EmailStruct) error {
	if es.From == "" {
		es.From = "kimi0230@gmail.com"
	}

	if es.Body == "" {
		es.Body = "Hi Kimi!"
	}

	if es.Subject == "" {
		es.Subject = "Hello Kimi!"
	}

	receivers := strings.Split(es.To, ",")
	mail := gomail.NewMessage()
	mail.SetHeader("From", es.From)
	mail.SetHeader("To", receivers...)
	mail.SetHeader("Subject", es.Subject)
	mail.SetBody("text/html", es.Body)
	// m.Attach("/home/Alex/lolcat.jpg")

	dialer := gomail.NewDialer(smtp, tlsPort, user, password)

	// Send the email
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("---------------- emailService:", err)
		return err
	}
	return nil
}
