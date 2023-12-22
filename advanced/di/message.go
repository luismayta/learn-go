package di

type IMessage interface {
	Value() string
}

type Message struct {
	v string
}

func (m Message) Value() string {
	return m.v
}

func NewMessage(phrase string) IMessage {
	return &Message{
		v: phrase,
	}
}
