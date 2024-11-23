package ws

import (
	"ChatServer/loggers"
	"ChatServer/ws/messagehandlers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
	"strings"
	"sync"
)

var hasInitialized = false
var upgrader = websocket.Upgrader{}

var pkgMutex = sync.Mutex{}
var closed = false

func initWebSocket(engine *gin.Engine) {

	if hasInitialized {
		return
	}

	pkgMutex.Lock()
	defer pkgMutex.Unlock()
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	hasInitialized = true
	engine.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			loggers.DefaultLogger.GlobalLogger.Error("Error upgrading to websocket", slog.String("error", err.Error()))
			return
		}

		defer func(conn *websocket.Conn) {
			if closed {
				return
			}

			err := conn.Close()
			if err != nil {
				loggers.DefaultLogger.GlobalLogger.Error("Failed to close websocket", slog.String("error", err.Error()))
			}
		}(conn)

		conn.SetCloseHandler(func(code int, text string) error {
			pkgMutex.Lock()
			closed = true
			defer pkgMutex.Unlock()
			return conn.Close()
		})

		for !closed {
			mt, message, err := conn.ReadMessage()
			if !processReadError(err) {
				continue
			}

			handlerMgr := messagehandlers.GetMessageHandlerManager()
			handler := handlerMgr.GetHandler(mt)
			if handler == nil {
				loggers.DefaultLogger.GlobalLogger.Error("Unknown message type")
				continue
			}

			receiveError := (*handler).Handle(mt, message, conn)
			processHandlerError(receiveError)
		}
	})
}

func processReadError(readError error) bool {
	return processHandlerError(readError)
}

func processHandlerError(receiveError error) bool {
	if receiveError == nil {
		return true
	}

	errInfo := receiveError.Error()
	if strings.Contains(errInfo, "use of closed network connection") {
		return false
	}

	if strings.Contains(errInfo, "close 1000 (normal)") {
		return false
	}

	loggers.DefaultLogger.GlobalLogger.Error("Failed to handle message")
	loggers.DefaultLogger.GlobalLogger.Error(receiveError.Error())
	return false
}

func InitWebSocket(engine *gin.Engine) {
	initWebSocket(engine)
}
