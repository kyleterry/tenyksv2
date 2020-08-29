package adapter

import (
	"context"

	"github.com/kyleterry/tenyks/pkg/message"
)

type Adapter interface {
	GetName() string
	GetType() AdapterType
	Dial(ctx context.Context) error
	Close(ctx context.Context) error
	SendAsync(ctx context.Context, msg message.Message) error
	RegisterMessageHandler(message.HandlerFunc)
}

type AdapterType int

const (
	AdapterTypeIRC AdapterType = iota
)

var AdapterTypeMapping = map[string]AdapterType{
	"irc": AdapterTypeIRC,
}
