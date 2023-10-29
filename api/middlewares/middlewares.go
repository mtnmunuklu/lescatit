package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/security"
)

// Authenticate provides the authentication process middleware.
func Authenticate(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString, err := security.ExtractToken(c)
		if err != nil {
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("error on parse token:", err.Error())
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		if !token.Valid {
			log.Println("invalid token:", tokenString)
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		return next(c)
	}
}
