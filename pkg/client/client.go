package client

import (
	"context"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
	"github.com/kyleterry/tenyks/pkg/message"
	"github.com/prometheus/common/log"
)

type Client struct {
	addr string
	conn *websocket.Conn
}

func (c *Client) Dial(ctx context.Context) error {
	u := &url.URL{Scheme: "ws", Host: c.addr, Path: "/"}

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return err
	}

	c.conn = conn

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Error(err)
		}

		pretty.Println(string(message))
	}

	return nil
}

func (c *Client) ReadMessage(ctx context.Context) (message.Message, error) {
	// _, message, err := c.conn.ReadMessage()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read message: %w", err)
	// }

	// return msg, nil
	return nil, nil
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}
