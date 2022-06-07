package util

import (
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	gomail "gopkg.in/mail.v2"
	"log"
	"net/smtp"
	"time"
)

func Sender() {
	// Sender data.
	from := "info@timtube.org"
	password := "eNJ8@1n3$A"

	// Receiver email address.
	to := []string{
		"216093805@mycput.ac.za",
	}

	// smtp server configuration.
	smtpHost := "smtp.ionos.com"
	smtpPort := "587"

	now := time.Now()
	fmt.Println(now.Format("EEE, d MMM yyyy HH:mm:ss z"))
	// Message.
	message := []byte("From: info@timtube.org\r\n" +
		"To: 216093805@mycput.ac.za\r\n" +
		"Subject: Golang testing mail\r\n" +
		"Welcome to Go! kjhdsjfhd kjdkfjjkjdd dkjskd\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
func Sender2() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "info@timtube.org")

	// Set E-Mail receivers
	m.SetHeader("To", "espoirditekemena@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.ionos.com", 587, "info@timtube.org", "eNJ8@1n3$A")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return

}

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func sender3() {
	server := mail.NewSMTPClient()
	server.Host = "smtp.ionos.com"
	server.Port = 587
	server.Username = "info@timtube.org"
	server.Password = "eNJ8@1n3$A"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("info@timtube.org")
	email.AddTo("espoirditekemena@gmail.com")
	//email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	//email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}
