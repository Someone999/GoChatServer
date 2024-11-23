package messagehandlers

import (
	"errors"
	"github.com/gorilla/websocket"
)

type PongMessageHandler struct {
}

func (p *PongMessageHandler) Handle(msgType int, _ []byte, _ *websocket.Conn) error {
	if msgType != websocket.PongMessage {
		return errors.New("unexpected message type, message type should be websocket.PingMessage")
	}

	return nil
}
