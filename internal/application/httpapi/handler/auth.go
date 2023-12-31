package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/components"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/views"
	"github.com/tomazcx/go-chat-app/internal/application/httpapi/utils"
	"github.com/tomazcx/go-chat-app/internal/data/user"
)

type AuthHandler struct{}

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

func (a *AuthHandler) HandleLogin(c *fiber.Ctx) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	validateUserUseCase := user.NewValidateUserUseCase()
	userValid := validateUserUseCase.Execute(login, password)

	if !userValid {
		c.Set("Content-Type", "text/html")
		return components.ErrorMessage("Invalid credentials").Render(c.Context(), c.Response().BodyWriter())
	}

	utils.SetUserSession(c, login)

	c.Response().Header.Add("HX-Redirect", "/")
	return c.Status(http.StatusSeeOther).SendString("")
}

func (a *AuthHandler) HandleLogout(c *fiber.Ctx) error {
	sess, _ := utils.GetUserSession(c)
	sess.Destroy()
	return c.Redirect("/auth")
}

func (a *AuthHandler) HandleRegister(c *fiber.Ctx) error {

	login := c.FormValue("login")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")

	if password != confirmPassword {
		c.Set("Content-Type", "text/html")
		return components.ErrorMessage("Passwords must match").Render(c.Context(), c.Response().BodyWriter())
	}

	dto := user.CreateUserDTO{Login: login, Password: password, ConfirmPassword: confirmPassword}

	createUserUseCase := user.NewCreateUserUseCase()
	user, err := createUserUseCase.Execute(dto)

	if err != nil {
		c.Set("Content-Type", "text/html")
		return components.ErrorMessage(err.Error()).Render(c.Context(), c.Response().BodyWriter())
	}

	utils.SetUserSession(c, user.Login)

	c.Response().Header.Add("HX-Redirect", "/")
	return c.Status(http.StatusSeeOther).SendString("")
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}
