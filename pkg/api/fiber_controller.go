package api

import "github.com/gofiber/fiber/v2"

type FiberController interface {
	RegisterRoutes(router fiber.Router)
}
type FiberApiController interface {
	RegisterRoutes(router fiber.Router)
}
