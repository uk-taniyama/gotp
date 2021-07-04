package controllers

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func onConnect(s socketio.Conn) error {
	s.SetContext("")
	logger.Info("connected:", s.Namespace(), s.ID())
	s.Join("X")
	return nil
}

func onError(s socketio.Conn, e error) {
	logger.Info("error:", e)
}

func onDisconnect(s socketio.Conn, reason string) {
	logger.Info("disconnect:", reason)
}

func onEventNotice(s socketio.Conn, msg string) {
	logger.Info("notice:", msg)
	s.Emit("reply", "have "+msg)
	server.BroadcastToRoom("", "X", "reply", "have "+msg)
}
func onEventChatMsg(s socketio.Conn, msg string) string {
	s.SetContext(msg)
	s.Emit("reply", msg)
	return "recv " + msg
}

func onEventBye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()
	return last
}

func NewSocketIOServer() *socketio.Server {
	server = socketio.NewServer(nil)
	server.OnConnect("/", onConnect)
	server.OnError("/", onError)
	server.OnDisconnect("/", onDisconnect)
	server.OnEvent("/chat", "msg", onEventChatMsg)
	server.OnEvent("/", "notice", onEventNotice)
	server.OnEvent("/", "bye", onEventBye)
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	// defer server.Close()
	return server
}
