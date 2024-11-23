package protomessagehandler

import (
	"ChatServer/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type ProtoBasePacketHandler struct {
}

func (handler ProtoBasePacketHandler) Handle(data []byte, conn *websocket.Conn) error {
	var basePacket = generated.BasePacket{}
	err := proto.Unmarshal(data, &basePacket)
	if err != nil {
		return err
	}
	var msgType = basePacket.GetMessageType()

	if msgType == generated.MessageType_MESSAGE {
		h := MessageProtoMessageHandler{}
		return h.Handle(basePacket.GetMessage(), conn)
	}

	return nil
}
