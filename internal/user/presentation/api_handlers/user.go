package user_api_handlers

import (
	user_commands "github.com/James1752/gonertia-test/internal/user/application/commands"
	user_api_dto "github.com/James1752/gonertia-test/internal/user/presentation/dto/api"
	api "github.com/James1752/gonertia-test/pkg/api"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type UserApiHandler struct {
}

func NewUserApiHandler() *UserApiHandler {
	return &UserApiHandler{}
}

func (uc *UserApiHandler) RegisterRoutes(router fiber.Router) {
	userRouter := router.Group("/user")

	//Register
	userRouter.Post("/register", api.NewFiberRequestHandler(func(c *fiber.Ctx, v *validator.Validate) (uuid.UUID, error) {
		//Parse body into struct
		body := &user_api_dto.UserRequestDto{}
		if err := c.BodyParser(body); err != nil {
			return uuid.Nil, err
		}

		// Validate body
		if err := v.Struct(body); err != nil {
			return uuid.Nil, err
		}

		//Cosntruct and send command
		command := user_commands.NewRegisterUserCommand(body.FirstName, body.LastName, body.Email)
		if result, err := mediatr.Send[*user_commands.RegisterUserCommand, uuid.UUID](c.Context(), command); err != nil {
			return uuid.Nil, err
		} else {
			return result, nil
		}
	}).OnSuccess(func(c *fiber.Ctx, res uuid.UUID) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"userId": res,
		})
	}).Execute)
}
