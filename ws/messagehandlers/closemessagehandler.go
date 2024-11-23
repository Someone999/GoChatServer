package messagehandlers

import "github.com/gorilla/websocket"

type CloseMessageHandler struct {
}

func (CloseMessageHandler CloseMessageHandler) Handle(_ int, _ []byte, _ *websocket.Conn) error {
	return nil
}
