package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	host           string
	port           string
	username       string
	password       string
	frontendAppUrl string
	isInitialized  = false
)

func initialize() {
	if isInitialized {
		return
	}

	host = os.Getenv("SMTP_HOST")
	port = os.Getenv("SMTP_PORT")
	username = os.Getenv("SMTP_USERNAME")
	password = os.Getenv("SMTP_PASSWORD")
	frontendAppUrl = os.Getenv("FRONTEND_APP_URL")

	isInitialized = true
}

func SendEmail(from string, to []string, subject string, message string) error {
	initialize()

	address := fmt.Sprintf("%s:%s", host, port)

	auth := smtp.PlainAuth("", username, password, host)

	_from := fmt.Sprintf("From: %s\r\n", from)
	_to := fmt.Sprintf("To: %s\r\n", strings.Join(to, ","))
	_subject := fmt.Sprintf("Subject: %s\r\n\r\n", subject)

	msg := []byte(_from + _to + _subject + message)

	return smtp.SendMail(address, auth, from, to, msg)
}

func SendInvitationEmail(to string, token string) error {
	initialize()

	subject := "[Learning App] you've been inviteted to access the app"

	activationUrl := frontendAppUrl + "/activate?token=" + token

	var err = SendEmail(username, []string{to}, subject, activationUrl)

	return err
}
