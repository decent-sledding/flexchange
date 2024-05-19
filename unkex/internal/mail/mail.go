package mail_api

import (
	"log"
	"errors"

	"unkex/config"
)


logger := log.New(os.Stdin)


type Message struct {
	From 	string
	To   	string
	Subject string
	Message [...]byte
}


func sendEmail(msg *Message) error {
	conf := config.GetConfig().Mail
	auth := smtp.PlainAuth("", msg.From, conf.SmtpPassword, conf.SmtpHost)
	fullHost := conf.SmtpHost + ':' + conf.SmtpPort

	if smtp.SendMail(fullHost, auth, msg.From, msg.To, msg.Message) != nil {
		log.Printf("Could not send e-mail")
		return errors.New("Email sending error")
	}

	fmt.Println("Email Sent Successfully!")
}

