package tools

import (
	"errors"
	"net/smtp"
	"os"

	"github.com/hermanowicz/todos/defs"
)

// maybe https://github.com/wneessen/go-mail/wiki/Getting-started

type SMTP_Cruds struct {
	HostName string
	AuthData smtp.Auth
}

func NewSmtpCruds() (*SMTP_Cruds, error) {
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_user := os.Getenv("SMTP_USER")
	smtp_pass := os.Getenv("SMTP_PASSWORD")

	if smtp_host == "" || smtp_pass == "" || smtp_user == "" {
		return nil, errors.New("one of the SMTP Server login data points in Env was empty, please check it")
	}

	return &SMTP_Cruds{
		HostName: smtp_host,
		AuthData: smtp.PlainAuth("", smtp_user, smtp_pass, smtp_host),
	}, nil
}

func SMTP_SendMial(mBody defs.MailMessage, crud *SMTP_Cruds) error {
	err := smtp.SendMail(string(crud.HostName)+":465", crud.AuthData, mBody.From, mBody.Recipients, mBody.Msg)
	if err != nil {
		return err
	}
	return nil
}
