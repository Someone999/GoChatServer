package main

import (
	"ChatServer/ws"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var engine = gin.Default()
	ws.InitWebSocket(engine)
	runErr := engine.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
		return
	}
}
