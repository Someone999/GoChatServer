package protomessagehandler

import (
	messagehandlers "awesomeProject/objectmanager"
	"sync"
)

var internalProtoMessageHandlerManager messagehandlers.HandlerManager[ProtoMessageHandler]

var pkgMutex = sync.Mutex{}

func GetProtoMessageHandlerManager() *messagehandlers.HandlerManager[ProtoMessageHandler] {
	pkgMutex.Lock()
	defer pkgMutex.Unlock()

	return &internalProtoMessageHandlerManager
}
