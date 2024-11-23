package messagehandlers

import (
	"ChatServer/protomessagehandler"
	"github.com/gorilla/websocket"
)

type BinaryMessageHandler struct {
}

func (BinaryMessageHandler) Handle(_ int, message []byte, conn *websocket.Conn) error {
	//TODO: 实现proto计息和处理
	h := protomessagehandler.ProtoBasePacketHandler{}
	return h.Handle(message, conn)
	//return conn.WriteMessage(websocket.TextMessage, []byte("Proto handler is not implemented now."))
}
