package websocket

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/components"
)

type SocketMessage struct {
	Author  string `json:"author"`
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
		server.conns[c] = true
		for {
			mt, payload, err := c.ReadMessage()
			if err != nil {
				log.Println("connection closed.")
				delete(server.conns, c)
				break

			}

			msg := SocketMessage{}
			if err = json.Unmarshal(payload, &msg); err != nil {
				log.Printf("invalid message format: %v\n", err)
				continue
			}

			if len(msg.Author) == 0 {
				msg.Author = "Anonymous"
			}

			buffer := &bytes.Buffer{}
			components.Message(msg.Author, msg.Message).Render(context.Background(), buffer)
			server.Broadcast(buffer.Bytes(), mt)
		}
	}))
}
