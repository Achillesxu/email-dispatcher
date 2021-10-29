package mail

type Mail struct {
	smtp     string
	user     string
	password string
	port     uint
}

func NewMail(smtp, user, password string, port uint) *Mail {
	return &Mail{
		smtp:     smtp,
		user:     user,
		password: password,
		port:     port,
	}
}

func (m Mail) Smtp() string {
	return m.smtp
}

func (m Mail) User() string {
	return m.user
}

func (m Mail) Password() string {
	return m.password
}

func (m Mail) Port() uint {
	return m.port
}
