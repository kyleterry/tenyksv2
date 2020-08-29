package service

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
	"github.com/kyleterry/tenyks/pkg/adapter"
	"github.com/kyleterry/tenyks/pkg/message"
	"github.com/kyleterry/tenyks/pkg/registry"
	"github.com/prometheus/common/log"
)

type WebsocketServer struct {
	upgrader     *websocket.Upgrader
	out          chan message.Message
	baseRegistry registry.AdapterRegistry
}

func (ws *WebsocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer conn.Close()

	readCh := make(chan *message.ChatMessage)
	go func() {
		for {
			msg := &message.ChatMessage{}
			if err := conn.ReadJSON(msg); err != nil {
				log.Error(err)
				continue
			}

			readCh <- msg
		}
	}()

	for {
		select {
		case msg := <-ws.out:
			pretty.Println("got", msg)
			if err := conn.WriteJSON(msg); err != nil {
				log.Error(err)
			}
		case msg := <-readCh:
			// TODO find path
			for _, adapters := range ws.baseRegistry.GetAdaptersFor {
				c.SendAsync(r.Context(), msg)
			}
		case <-r.Context().Done():
			break
		}
	}
}

func (ws *WebsocketServer) RegisterHandler() {
	for _, c := range ws.conns {
		c.RegisterMessageHandler(func(msg message.Message) {
			ws.out <- msg
		})
	}
}

func NewWebsocketServer(ar adapter.Registry) *WebsocketServer {
	return &WebsocketServer{
		upgrader:     &websocket.Upgrader{},
		out:          make(chan message.Message, 10),
		baseRegistry: ar,
	}
}
