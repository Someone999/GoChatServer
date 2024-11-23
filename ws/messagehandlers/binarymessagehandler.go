package messagehandlers

import (
	"awesomeProject/loggers"
	"github.com/gorilla/websocket"
)

type BinaryMessageHandler struct {
}

func (BinaryMessageHandler) Handle(_ int, message []byte, conn *websocket.Conn) error {
	loggers.DefaultLogger.GlobalLogger.Error("Unexpected binary message.")
	return conn.WriteMessage(websocket.TextMessage, []byte("Binary message is not supported now."))
}
