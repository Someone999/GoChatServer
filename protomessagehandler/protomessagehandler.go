package protomessagehandler

import "github.com/gorilla/websocket"

type ProtoMessageHandler interface {
	Handle(data []byte, conn *websocket.Conn)
}
