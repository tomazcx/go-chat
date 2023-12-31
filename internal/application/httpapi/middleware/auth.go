package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {

	session, err := utils.GetUserSession(c)

	if err != nil {
		return c.Redirect("/auth")
	}

	user := session.Get("user")

	if user == nil {
		return c.Redirect("/auth")
	}

	return c.Next()
}
