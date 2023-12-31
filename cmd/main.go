package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/routes"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/utils"
	"github.com/tomazcx/go-chat-app/internal/application/websocket"
	"github.com/tomazcx/go-chat-app/internal/infra/database"
)


func main(){

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading env variables: %v", err)
	}

	app := routes.MakeRouter()
	websocket.Init(app)
	err := database.InitializaDB()
	utils.InitSessionStore()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))

}
