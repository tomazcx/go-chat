package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/views"
)

type ChatHandler struct{}

func (h *ChatHandler) Index(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return views.Chat().Render(c.Context(), c.Response().BodyWriter())
}

func NewChatHandler() *ChatHandler{
	return &ChatHandler{}
}
