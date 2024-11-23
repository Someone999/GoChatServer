package messagehandlers

import (
	messagehandlers "ChatServer/objectmanager"
	"github.com/gorilla/websocket"
	"sync"
)

var pkgMutex = sync.Mutex{}

var internalMessageHandlerManager *messagehandlers.HandlerManager[MessageHandler] = nil

func GetMessageHandlerManager() *messagehandlers.HandlerManager[MessageHandler] {
	pkgMutex.Lock()
	defer pkgMutex.Unlock()

	if internalMessageHandlerManager == nil {
		internalMessageHandlerManager = messagehandlers.NewHandlerManager[MessageHandler]()
		val := *internalMessageHandlerManager
		val.AddHandler(websocket.PingMessage, &PingMessageHandler{})
		val.AddHandler(websocket.PongMessage, &PongMessageHandler{})
		val.AddHandler(websocket.CloseMessage, &CloseMessageHandler{})
		val.AddHandler(websocket.BinaryMessage, &BinaryMessageHandler{})
		val.AddHandler(websocket.TextMessage, &TextMessageHandler{})
	}

	return internalMessageHandlerManager
}
