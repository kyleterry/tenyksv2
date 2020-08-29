package message

import "io"

type MessageType string

const (
	MessageTypeChat    MessageType = "chat"
	MessageTypeControl MessageType = "control"
)

// Message can encode, decode and validate messages flowing through tenkys
type Message interface {
	Encode(w io.Writer) error
	Decode(r io.Reader) error
	Validator() Validator
}

// Validator will validate a byte slice to see if it's the right type of
// message to decode
type Validator interface {
	Validate(b []byte) error
}

// Transformer transforms messages to and from Message and is used to communicate
// between a chat protocol and a service connected to Tenyks via a websocket.
type Transformer interface {
	ToTenkysMessage() (Message, error)
	FromTenyksMessage(Message) error
}
