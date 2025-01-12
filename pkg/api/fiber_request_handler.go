package api

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Fn[T any] func(context *fiber.Ctx, validator *validator.Validate) (T, error)
type OnSuccessFn[T any] func(context *fiber.Ctx, data T)
type OnFailureFn func(*fiber.Ctx, error)

type FiberRequestHandler[T any] struct {
	validator *validator.Validate
	fn        Fn[T]
	onSuccess OnSuccessFn[T]
	onFailure OnFailureFn
}

func NewFiberRequestHandler[T any](fn Fn[T]) *FiberRequestHandler[T] {
	return &FiberRequestHandler[T]{
		validator: validator.New(),
		fn:        fn,
		onSuccess: func(context *fiber.Ctx, data T) {
			context.SendStatus(fiber.StatusOK)
		},
		onFailure: func(context *fiber.Ctx, err error) {
			context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Controller Failure",
				"details": err.Error(),
			})
		},
	}
}

func (c *FiberRequestHandler[T]) Execute(context *fiber.Ctx) error {
	result, err := c.fn(context, c.validator)

	if err != nil {
		c.onFailure(context, err)
		return err
	}

	c.onSuccess(context, result)

	return nil
}

func (c *FiberRequestHandler[T]) OnSuccess(fn OnSuccessFn[T]) *FiberRequestHandler[T] {
	c.onSuccess = fn

	return c
}

func (c *FiberRequestHandler[T]) OnFailure(fn OnFailureFn) *FiberRequestHandler[T] {
	c.onFailure = fn

	return c
}
