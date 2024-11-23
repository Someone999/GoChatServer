package messagehandlers

import "github.com/gorilla/websocket"

type CloseMessageHandler struct {
}

func (CloseMessageHandler CloseMessageHandler) Handle(msgType int, data []byte, conn *websocket.Conn) error {
	return nil
}
