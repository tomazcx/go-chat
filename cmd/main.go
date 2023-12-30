package main

import (
	"log"

	"github.com/tomazcx/go-chat-app/internal/application/httpapi/routes"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/utils"
	"github.com/tomazcx/go-chat-app/internal/application/websocket"
	"github.com/tomazcx/go-chat-app/internal/infra/database"
)


func main(){
	app := routes.MakeRouter()
	websocket.Init(app)
	err := database.InitializaDB()
	utils.InitSessionStore()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Fatal(app.Listen(":8000"))

}
