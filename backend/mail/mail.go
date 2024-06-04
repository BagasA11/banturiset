package mail

import (
	"net/smtp"

	"github.com/bagasa11/banturiset/helpers"

	"fmt"
	"os"
)

func Notify(subject string, to string, port uint) error {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASS"), os.Getenv("SMTP"))
	bodymail := []byte(helpers.BodyF(subject))
	return smtp.SendMail(fmt.Sprintf("%s:%d", os.Getenv("SMTP"), port), auth, os.Getenv("MAIL_USERNAME"), []string{to}, bodymail)
}
