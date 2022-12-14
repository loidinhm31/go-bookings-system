package main

import (
	"fmt"
	"github.com/loidinhm31/go-bookings-system/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
	"strings"
	"time"
)

func listenForMail() {
	log.Println("Starting mail listener...")
	go func() {
		for {
			msg := <-app.MailChannel
			sendMessage(msg)
		}
	}()
}

func sendMessage(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "127.0.0.1"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.TemplateMail == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", m.TemplateMail))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent")
	}
}
