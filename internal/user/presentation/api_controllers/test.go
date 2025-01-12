package user_controllers

import (
	api "github.com/James1752/gonertia-test/pkg/api"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type TestApiController struct {
}

func NewTestApiController() *TestApiController {
	return &TestApiController{}
}

func (uc *TestApiController) RegisterRoutes(router fiber.Router) {
	userRouter := router.Group("/test")

	//Register
	userRouter.Get("/test", api.NewFiberRequestHandler(func(c *fiber.Ctx, v *validator.Validate) (any, error) {
		return nil, nil
	}).OnSuccess(func(c *fiber.Ctx, res any) {
		c.Status(fiber.StatusOK).SendString("Hello from test")
	}).Execute)
}
