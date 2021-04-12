package kindle_emailer

import (
	"fmt"
	"kindle_clipping_exporter/kindle"
	"log"
	"net/smtp"
	"strings"
)

const messageHeader = "To: %s\r\nSubject: Your most recent kindle clippings\r\n"

type Credentials struct {
	FromEmail         string
	FromEmailPassword string
	ToEmail           string
}

func SendEmail(d kindle.Device, cred *Credentials) {
	if len(d.NewClippings) == 0 {
		log.Println("No new clippings, skipping email.")
		return
	}
	// Sender data.
	from := cred.FromEmail
	password := cred.FromEmailPassword

	// Receiver email address.
	to := []string{
		cred.ToEmail,
	}

	// smtp server configuration for gmail.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	var msg strings.Builder
	msg.Write([]byte(fmt.Sprintf(messageHeader, cred.ToEmail)))
	for _, clipping := range d.NewClippings {
		msg.WriteString(clipping.ToString())
		msg.WriteString("\n==========\n")
	}
	msg.WriteString("\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg.String()))
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Email Sent Successfully!")
}
