package main

import (
	"Golearn/NetProgram/Websocket\ tutorial/core"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Handle(ctx *gin.Context) {
	conn, _ := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	wsConn, _ := core.NewConnection(conn)
	wsConn.Start()
	go func() {
		for {
			time.Sleep(time.Second)
			wsConn.WriteMessage([]byte("server"))
		}
	}()
	for {
		msg, _ := wsConn.ReadMessage()
		wsConn.WriteMessage(msg)
	}
}

func main() {
	router := gin.Default()
	router.GET("/ws", Handle)
	router.Run(":8989")
}
