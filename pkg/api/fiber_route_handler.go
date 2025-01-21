package api

import "github.com/gofiber/fiber/v2"

type FiberHandler interface {
	RegisterRoutes(router fiber.Router)
}
