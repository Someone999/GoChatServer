package messagehandlers

import (
	"github.com/gorilla/websocket"
)

type BinaryMessageHandler struct {
}

func (BinaryMessageHandler) Handle(_ int, message []byte, conn *websocket.Conn) error {
	//TODO: 实现proto计息和处理
	return conn.WriteMessage(websocket.TextMessage, []byte("Proto handler is not implemented now."))
}
