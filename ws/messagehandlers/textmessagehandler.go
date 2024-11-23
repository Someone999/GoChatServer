package messagehandlers

import (
	"github.com/gorilla/websocket"
)

type TextMessageHandler struct {
}

func (TextMessageHandler TextMessageHandler) Handle(_ int, _ []byte, _ *websocket.Conn) error {
	return nil
}
