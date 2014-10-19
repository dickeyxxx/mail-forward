package main

import (
	"fmt"
	"net/smtp"
	"net/textproto"
)

func SendMail(from string, to []string, lines []string) error {
	c, err := smtp.Dial("alt1.gmail-smtp-in.l.google.com:25")
	defer c.Close()
	if err != nil {
		return err
	}
	fmt.Println("Sending email from", from)
	if err := c.Mail(from); err != nil {
		return err
	}
	for _, to := range to {
		to = "dickeyxxx@gmail.com"
		fmt.Println("Sending email to", to)
		if err := c.Rcpt(to); err != nil {
			return err
		}
	}
	wc, err := c.Data()
	if err != nil {
		return err
	}
	for _, line := range lines {
		if _, err = fmt.Fprintf(wc, line); err != nil {
			return err
		}
	}
	wc.Close()
	err = closeWithFullError(c)
	return c.Quit()
}

func closeWithFullError(c *smtp.Client) error {
	code, message, err := c.Text.ReadResponse(0)
	if err != nil {
		return err
	}
	if code != 250 {
		return &textproto.Error{Code: code, Msg: message}
	}
	return nil
}
