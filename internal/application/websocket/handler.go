package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/websocket/socketutils"
)

type SocketMessage struct {
	Message string `json:"message"`
}

func Init(app *fiber.App) {
	server := NewServer()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		m := sync.Mutex{}
		m.Lock()
		server.conns[c] = true
		m.Unlock()
		userName := c.Query("userName")

		//notifies that the a user logged in
		bufferizedMessage := socketutils.BufferizeMessage(fmt.Sprintf("The user %s has entered the chat!", userName), "")
		server.Broadcast(bufferizedMessage, 1)

		for {
			mt, payload, err := c.ReadMessage()
			if err != nil {
				log.Println("connection closed.")
				delete(server.conns, c)
				bufferizedMessage := socketutils.BufferizeMessage(fmt.Sprintf("%s has left the chat!", userName), "")
				server.Broadcast(bufferizedMessage, 1)
				break
			}

			msg := SocketMessage{}
			if err = json.Unmarshal(payload, &msg); err != nil {
				log.Printf("invalid message format: %v\n", err)
				continue
			}
			bufferizedMessage := socketutils.BufferizeMessage(msg.Message, userName)
			server.Broadcast(bufferizedMessage, mt)
		}
	}))
}
