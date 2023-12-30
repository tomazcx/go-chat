package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi"
	"github.com/tomazcx/go-chat-app/internal/application/websocket"
)


func main(){

	app := fiber.New()

	httpapi.DefineRoutes(app)
	websocket.Init(app)

	log.Fatal(app.Listen(":8000"))

}
