package messagehandlers

import (
	"github.com/gorilla/websocket"
)

type TextMessageHandler struct {
}

func (TextMessageHandler TextMessageHandler) Handle(msgType int, data []byte, conn *websocket.Conn) error {
	return nil
}
