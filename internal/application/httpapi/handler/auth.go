package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/components"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/views"
)

type AuthHandler struct {}

func (a *AuthHandler) Index(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return views.Auth().Render(c.Context(), c.Response().BodyWriter())
}

func (a *AuthHandler) LoginForm(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return components.LoginForm().Render(c.Context(), c.Response().BodyWriter())
}

func (a *AuthHandler) RegisterForm(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return components.RegisterForm().Render(c.Context(), c.Response().BodyWriter())
}

func NewAuthHandler() *AuthHandler{
	return &AuthHandler{}
}
