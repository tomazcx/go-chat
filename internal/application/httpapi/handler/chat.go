package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/views"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/utils"
)

type ChatHandler struct{}

func (h *ChatHandler) Index(c *fiber.Ctx) error {
	sess, _ := utils.GetUserSession(c)
	userName := sess.Get("user").(string)
	c.Set("Content-Type", "text/html")
	return views.Chat(userName).Render(c.Context(), c.Response().BodyWriter())
}

func NewChatHandler() *ChatHandler{
	return &ChatHandler{}
}
