package messagehandlers

import "sync"

var pkgMutex = sync.Mutex{}

type HandlerManager[T any] struct {
	handlers map[int]*T
}

func NewHandlerManager[T any]() *HandlerManager[T] {
	return &HandlerManager[T]{handlers: make(map[int]*T)}
}

func (handlerManager *HandlerManager[T]) AddHandler(msgType int, handler T) {
	pkgMutex.Lock()
	defer pkgMutex.Unlock()
	handlerManager.handlers[msgType] = &handler

}

func (handlerManager *HandlerManager[T]) RemoveHandler(msgType int) {
	pkgMutex.Lock()
	defer pkgMutex.Unlock()
	delete(handlerManager.handlers, msgType)

}

func (handlerManager *HandlerManager[T]) GetHandler(msgType int) *T {
	val, exists := handlerManager.handlers[msgType]
	if !exists {
		return nil
	}

	return val
}
