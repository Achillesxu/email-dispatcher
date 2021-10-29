package config

type MailService struct {
	smtp     string
	user     string
	password string
	port     uint
}

func NewMailService() *MailService {
	return &MailService{
		smtp:     GetConfig().MailSmtp,
		user:     GetConfig().MailUser,
		password: GetConfig().MailPassword,
		port:     GetConfig().MailPort,
	}
}
