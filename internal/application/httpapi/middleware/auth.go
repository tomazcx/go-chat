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

	user, ok := session.Get("user").(string)

	if len(user) == 0 || !ok {
		return c.Redirect("/auth")
	}

	return c.Next()
}
