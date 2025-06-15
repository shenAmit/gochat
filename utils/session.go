package utils

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New(session.Config{
	CookieName:     "admin_session",
	Expiration:     24 * time.Hour,
	CookieHTTPOnly: true,
})
