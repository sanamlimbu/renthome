package email

import (
	"fmt"
	"net/smtp"

	"github.com/ninja-software/terror/v2"
)

type Mailer struct {
	Username string
	Password string
	Host     string
	Port     string
	Auth     *smtp.Auth
}

func NewMailer(username, password, host, port string) *Mailer {

	auth := smtp.PlainAuth("", username, password, host)

	result := &Mailer{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Auth:     &auth,
	}

	return result
}

func (mailer *Mailer) SendEmail(from string, to []string, message string) error {
	hostPort := fmt.Sprintf("%s:%s", mailer.Host, mailer.Port)

	err := smtp.SendMail(hostPort, *mailer.Auth, from, to, []byte("Hello World"))
	if err != nil {
		return terror.Error(err, "cannot send email")
	}

	return nil
}
