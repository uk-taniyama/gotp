package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Header)
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("WebSocket Upgrade", err)
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		logger.Info("CONN", msg)
		conn.WriteMessage(t, msg)
	}
}

func GetWS(c *gin.Context) {
	handler(c.Writer, c.Request)
}
