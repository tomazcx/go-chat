package websocket

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App){
	app.Use("/ws", func(c *fiber.Ctx)error {
		if websocket.IsWebSocketUpgrade(c){
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn){

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				break;

			}

			log.Printf("received: %s", msg)

			if err := c.WriteMessage(mt, msg); err != nil {
				log.Println("error writting: ", err)
				break
			}
		}
	}))	
}
