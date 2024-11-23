package ws

import (
	"awesomeProject/loggers"
	"awesomeProject/ws/messagehandlers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
	"sync"
)

var hasInitialized = false
var upgrader = websocket.Upgrader{}

var pkgMutex = sync.Mutex{}

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
			if err != nil {
				loggers.DefaultLogger.GlobalLogger.Error("Error upgrading to websocket", slog.String("error", err.Error()))
			}
			return
		}

		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				loggers.DefaultLogger.GlobalLogger.Error("Failed to close websocket", slog.String("error", err.Error()))
			}
		}(conn)

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				continue
			}

			if mt == websocket.PingMessage {
				_ = conn.WriteMessage(websocket.PongMessage, []byte{})
				return
			}

			handlerMgr := messagehandlers.GetMessageHandlerManager()
			handler := handlerMgr.GetHandler(mt)
			if handler == nil {
				loggers.DefaultLogger.GlobalLogger.Error("Unknown message type")
				continue
			}

			receiveError := (*handler).Handle(mt, message, conn)
			if receiveError != nil {
				loggers.DefaultLogger.GlobalLogger.Error("Failed to handle message")
			}
		}
	})
}

func InitWebSocket(engine *gin.Engine) {
	initWebSocket(engine)
}
