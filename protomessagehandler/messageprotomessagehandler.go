package protomessagehandler

import (
	"ChatServer/generated/packet/message"
	"fmt"
	"github.com/gorilla/websocket"
)

type MessageProtoMessageHandler struct {
}

func (handler MessageProtoMessageHandler) Handle(data *message.Message, _ *websocket.Conn) error {
	fmt.Println(string(data.Data))
	return nil
}
