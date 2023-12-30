package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/handler"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/middleware"
)

type ChatRouter struct {
	h *handler.ChatHandler
}

func (c *ChatRouter) Load(r *fiber.App){
	r.Use("/", middleware.AuthMiddleware)
	r.Get("/", c.h.Index)
}

func NewChatRouter() *ChatRouter{
	return &ChatRouter{
		h: handler.NewChatHandler(),
	}
}
