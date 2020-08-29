package message

type Registry interface {
	RegisterMessageHandler(HandlerFunc)
}

type HandlerFunc func(Message)
