package messagehandlers

import (
	"errors"
	"github.com/gorilla/websocket"
)

type PingMessageHandler struct {
}

func (p *PingMessageHandler) Handle(msgType int, _ []byte, conn *websocket.Conn) error {
	if msgType != websocket.PingMessage {
		return errors.New("unexpected message type, message type should be websocket.PingMessage")
	}

	return conn.WriteMessage(websocket.PongMessage, nil)
}
