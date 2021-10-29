package message

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

func (m *Message) From() string {
	return m.from
}

func (m *Message) To() string {
	return m.to
}

func (m *Message) Subject() string {
	return m.subject
}

func (m *Message) Body() string {
	return m.body
}
