package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Handle(ctx *gin.Context) {
	wsConn, _ := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	_, data, _ := wsConn.ReadMessage()
	wsConn.WriteMessage(websocket.TextMessage, data)
}

func main() {
	router := gin.Default()
	router.GET("/ws", Handle)
	router.Run(":8989")
}
