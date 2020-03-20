package email

import (
	"bytes"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/aklinker1/anime-skip-backend/internal/utils"
)

var password string

type Email struct {
	To           []string
	Subject      string
	Template     string
	TemplateData map[string]string
}

func init() {
	password = utils.EnvString("EMAIL_PASSWORD")
}

func (email Email) Send() error {
	from := "aaronklinker1@gmail.com"
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	body, err := parseTemplate(email.Template, email.TemplateData)
	if err != nil {
		return err
	}
	content := []string{
		"MIME-version: 1.0;",
		"Content-Type: text/html; charset=\"UTF-8\";",
		"From: " + from,
		"To: " + strings.Join(email.To, ","),
		"Subject: " + email.Subject,
		"",
		body,
	}

	if utils.EnvBool("DISABLE_EMAILS") {
		return nil
	}

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		email.To,
		[]byte(strings.Join(content, "\n")),
	)
}

func parseTemplate(name string, data interface{}) (string, error) {
	t, err := template.ParseFiles("web/" + name)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
