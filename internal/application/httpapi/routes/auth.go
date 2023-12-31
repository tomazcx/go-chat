package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/handler"
)

type AuthRouter struct {
	h *handler.AuthHandler
}

func (a *AuthRouter) Load(r *fiber.App){
	r.Get("/auth", a.h.Index)	
	r.Get("/auth/login", a.h.LoginForm)	
	r.Get("/auth/register", a.h.RegisterForm)	
	r.Get("/auth/logout", a.h.HandleLogout)	
	r.Post("/auth/login", a.h.HandleLogin)
	r.Post("/auth/register", a.h.HandleRegister)	
}

func NewAuthRouter() *AuthRouter{
	return &AuthRouter{
		h: handler.NewAuthHandler(),
	}
}
