package messagehandlers

import (
	"github.com/gorilla/websocket"
)

type MessageHandler interface {
	Handle(msgType int, data []byte, conn *websocket.Conn) error
}
