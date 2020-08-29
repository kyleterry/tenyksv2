package message

type MessageEnvelope struct {
	Type    MessageType `json:"type"`
	Message interface{} `json:"message"`
}

func NewMessageEnvelope(msg Message) *MessageEnvelope {
	mt := MessageTypeChat

	if _, ok := msg.(*ControlMessage); ok {
		mt = MessageTypeControl
	}

	return &MessageEnvelope{
		Type:    mt,
		Message: msg,
	}
}
