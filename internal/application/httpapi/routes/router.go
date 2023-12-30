package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MakeRouter() *fiber.App {
	r := fiber.New()

	chatRouter := NewChatRouter()
	authRouter := NewAuthRouter()

	authRouter.Load(r)
	chatRouter.Load(r)

	return r
}

