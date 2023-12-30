package websocket

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func (s *Server) Broadcast(payload []byte, messateType int) {
	for conn := range s.conns {
		go func(ws *websocket.Conn){
			if err := ws.WriteMessage(messateType, payload); err != nil {
				log.Println("error writting: ", err)
			}
		}(conn)
	}
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

