package middlewares

import (
	"api/types"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit(max int, expiration int) func(*fiber.Ctx) error {

	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(expiration) * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(types.ErrorResponse{Message: types.ERROR_RATE_LIMIT_EXCEEDED})
		},
		// Storage: myCustomStorage{}
	})

}
