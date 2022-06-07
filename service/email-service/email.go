package email_service

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func Sender() {
	from := mail.NewEmail("espoir", "216093805@mycput.ac.za")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("espoir ditekemena", "espoirditekemena@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("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"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
