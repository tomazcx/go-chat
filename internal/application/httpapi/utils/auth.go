package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func InitSessionStore(){
	store = session.New()
}

func SetUserSession(login string, c *fiber.Ctx) (*session.Session, error){
	sess, err := store.Get(c)

	if err != nil {
		return nil, err
	}

	sess.Set("user", login)
	sess.SetExpiry(time.Hour * 24)
	return sess, nil
}

func GetUserSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)

	if err != nil {
		return nil, err
	}

	return sess, nil
}

