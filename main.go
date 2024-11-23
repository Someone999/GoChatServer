package main

import (
	"awesomeProject/ws"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var engine *gin.Engine = gin.Default()
	ws.InitWebSocket(engine)
	runErr := engine.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
		return
	}

}
