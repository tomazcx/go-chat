package main

import (
	"net/http"

	"github.com/tomazcx/go-chat-app/internal/application/handler"
)

func main(){
	chatHandler := handler.ChatHandler{}
	http.HandleFunc("/", chatHandler.Index)

	http.ListenAndServe(":8000", nil)	
}
