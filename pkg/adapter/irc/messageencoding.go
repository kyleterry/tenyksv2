package irc

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/kyleterry/tenyks/pkg/message"
)

// MessageEncoder takes a Message object and returns a string value for that
// message
type MessageEncoder interface {
	Encode(*Message) (string, error)
}

// MessageDecoder takes a message as a string and parses it into a Message object
type MessageDecoder interface {
	Decode(string) (*Message, error)
}

// RawMessageEncoder builds a raw IRC message string from a Message object
type RawMessageEncoder struct{}

func (e *RawMessageEncoder) Encode(msg *Message) (string, error) {
	params := msg.Params

	if msg.Trail != "" {
		params = append(params, fmt.Sprintf(":%s", msg.Trail))
	}

	msg.RawMsg = fmt.Sprintf("%s %s", msg.Command, strings.Join(params, " "))

	// TODO encode tags
	return fmt.Sprintf("%s\r\n", msg.RawMsg), nil
}

// NewRawMessageEncoder returns a new RawMessageEncoder object. It is used to
// encode Message objects into strings that can be sent to IRC servers.
func NewRawMessageEncoder() *RawMessageEncoder {
	return &RawMessageEncoder{}
}

// RawMessageDecoder takes a raw IRC message and parses it into a Message object
type RawMessageDecoder struct{}

func (d *RawMessageDecoder) Decode(s string) (*Message, error) {
	return ParseMessage(s)
}

// NewRawMessageDecoder returns a new RawMessageDecoder object. It is used to
// decode message strings sent by IRC servers into Message objects.
func NewRawMessageDecoder() *RawMessageDecoder {
	return &RawMessageDecoder{}
}

type tenyksChatMessageEncoder struct{}

func (tme *tenyksChatMessageEncoder) Encode(cmd *PrivmsgCommand) (message.Message, error) {
	tmsg := &message.ChatMessage{
		DestinationPath: "/",
		OriginPath:      "/",
		Content:         cmd.Message().Trail,
		Timestamp:       time.Now(),
	}

	return tmsg, nil
}

type tenyksChatMessageDecoder struct{}

func (tmd *tenyksChatMessageDecoder) Decode(msg message.Message) (*PrivmsgCommand, error) {
	var cmd *PrivmsgCommand

	switch msg.(type) {
	case *message.ChatMessage:
		theirs := msg.(*message.ChatMessage)

		_, target := path.Split(theirs.DestinationPath)

		cmd = NewPrivmsgCommand(target, theirs.Content)
	default:
		return nil, errors.New("unexpected message type")
	}

	return cmd, nil
}
