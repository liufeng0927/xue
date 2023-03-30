package email

import "time"

type Config struct {
	ExpireTime time.Duration
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string
	MailTitle  string
}
