package mail

import (
	"gopkg.in/gomail.v2"
)

var mailer *gomail.Message

type Config struct {
	Host         string
	Port         int
	Sender       string
	AuthEmail    string
	AuthPassword string
}

func init() {
	if mailer != nil {
		return
	}
	mailer = gomail.NewMessage()
}

func (c *Config) SetReceivers(emails ...string) {
	mailer.SetHeader("To", emails...)
}

func (c *Config) SetCC(address, name string) {
	mailer.SetAddressHeader("Cc", address, name)
}

func (c *Config) SetSubject(subject string) {
	mailer.SetHeader("Subject", subject)
}

func (c *Config) SetBody(body string) {
	mailer.SetBody("text/html", body)
}

func (c *Config) Send() error {

	mailer.SetHeader("From", c.AuthEmail)
	dialer := gomail.NewDialer(
		c.Host,
		c.Port,
		c.AuthEmail,
		c.AuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	mailer.Reset()

	return nil
}
