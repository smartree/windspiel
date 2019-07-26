package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/spf13/viper"
)

// TODO:
// - authenticated SMTP (see https://github.com/golang/go/wiki/SendingMail)
// - StartTLS
// - Mail template from configuration - needs some templating engine

var template = "This is windspiel on %s.\n\nSomebody connected to me on port %s from the IP %s at %s\n\nSent data was: %s\n\nYou may take a look. Stay safe!"

type mailEventHandler struct {
	sender            string
	receiverAddresses []string
	smtpServerAddress string
}

func (handler *mailEventHandler) configure(config *viper.Viper) error {
	log.Printf("Enabling mail events")
	registeredHandlers = append(registeredHandlers, handler)

	handler.sender = config.GetString("sender")
	handler.receiverAddresses = config.GetStringSlice("receivers")
	handler.smtpServerAddress = config.GetString("smtpServerAddress")

	// Test connection
	c, err := smtp.Dial(handler.smtpServerAddress)
	defer c.Close()

	return err
}

func (handler *mailEventHandler) processEvent(e event) {
	c, err := smtp.Dial(handler.smtpServerAddress)
	if err != nil {
		checkError(err, false)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail(handler.sender)

	for _, recipient := range handler.receiverAddresses {
		c.Rcpt(recipient)
	}
	// Send the email body.
	wc, err := c.Data()
	checkError(err, false)
	defer wc.Close()
	buf := bytes.NewBufferString(fmt.Sprintf(template,
		e.Target,
		e.Port,
		e.Src,
		e.Time.Format(time.RFC3339),
		e.Data))
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
