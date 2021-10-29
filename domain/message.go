package domain

type Message struct {
	from    string
	to      string
	subject string
	body    string
}

func NewMessage(from, to, subject, body string) Message {
	return Message{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}
