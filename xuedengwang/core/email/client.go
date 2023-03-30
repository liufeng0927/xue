package email

import (
	"gopkg.in/mail.v2"
)

type Client struct {
	ToEmail  string
	MailText string
	Config   Config
}

func NewClient(toEmail string, code string, config Config) *Client {
	mailText := "验证码：" + code + "此验证码用于更换邮箱绑定，请勿将验证码告知他人，有效期3分钟，请妥善保管。"
	return &Client{
		ToEmail:  toEmail,
		MailText: mailText,
		Config:   config,
	}
}
func (c *Client) Send() (bool, error) {

	m := mail.NewMessage()
	m.SetHeader("From", c.Config.SmtpEmail)
	m.SetHeader("To", c.ToEmail)
	m.SetHeader("Subject", c.Config.MailTitle)
	m.SetBody("text/html", c.MailText)
	d := mail.NewDialer(c.Config.SmtpHost, 465, c.Config.SmtpEmail, c.Config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	_ = d.DialAndSend(m)
	return true, nil

}
