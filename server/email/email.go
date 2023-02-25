package email

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
)

type Mailer struct {
	Username  string
	Password  string
	Host      string
	Port      string
	Templates map[string]*raymond.Template
}

func NewMailer(username, password, host, port string) (*Mailer, error) {

	result := &Mailer{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}

	// Read HTML templates and parse them
	templates := make(map[string]*raymond.Template)
	templateFiles, err := filepath.Glob("./email/templates/*.html")
	if err != nil {
		return nil, err
	}

	for _, templateFile := range templateFiles {
		templateName := filepath.Base(templateFile)
		templateData, err := os.ReadFile(templateFile)
		if err != nil {
			return nil, err
		}
		parsedTemplate, err := raymond.Parse(string(templateData))
		if err != nil {
			return nil, err
		}
		templates[templateName] = parsedTemplate
	}

	result.Templates = templates

	return result, nil
}

func (mailer *Mailer) SendAccountVerificationCode(reciever, name, code string) error {
	// Template key should match with HTML template file name
	templateKey := "reset_password.html"

	plainAuth := smtp.PlainAuth("", mailer.Username, mailer.Password, mailer.Host)
	hostPort := fmt.Sprintf("%s:%s", mailer.Host, mailer.Port)

	templateData := map[string]string{
		"code": code,
		"name": name,
	}

	htmlBody, err := mailer.Templates[templateKey].Exec(templateData)
	if err != nil {
		return err
	}

	from := mail.Address{Name: "renthome.com", Address: mailer.Username}

	// Convert the HTML body to bytes for use in the SMTP message
	messageBody := bytes.NewBufferString("")
	messageBody.WriteString("From: " + from.String() + "\r\n")
	messageBody.WriteString("To: " + reciever + "\r\n")
	messageBody.WriteString("Subject: " + "Account Verification Code" + "\r\n")
	messageBody.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")
	messageBody.WriteString(htmlBody)

	err = smtp.SendMail(hostPort, plainAuth, mailer.Username, []string{reciever}, messageBody.Bytes())
	if err != nil {
		return err
	}

	return nil
}
