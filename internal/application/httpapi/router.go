package httpapi


import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates"
)

func DefineRoutes(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error{
                c.Set("Content-Type", "text/html")
		return templates.Index().Render(c.Context(), c.Response().BodyWriter())
	})
}
