package controllers_user

import (
	"log"

	api_dto_user "github.com/James1752/gonertia-test/internal/api/dto/user"
	commmands_user "github.com/James1752/gonertia-test/internal/application/commands/user/register_user"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type UserController struct {
	validator *validator.Validate
}

func NewUserController() *UserController {
	router := &UserController{
		validator: validator.New(),
	}

	return router
}

func (controller *UserController) RegisterRoutes(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/user")
	userRouter.Post("/register", controller.registerUser)
}

// Register User
func (controller *UserController) registerUser(c *fiber.Ctx) error {
	body := &api_dto_user.UserRequestDto{}

	// Parse the body into the struct
	if err := c.BodyParser(body); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate the struct fields using the validator
	if err := controller.validator.Struct(body); err != nil {
		// If validation fails, return a 422 error with validation errors
		log.Printf("Validation error: %v\n", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	//Cosntruct and send command
	command := commmands_user.NewRegisterUserCommand(
		body.FirstName, body.LastName, body.Email,
	)
	result, err := mediatr.Send[*commmands_user.RegisterUserCommand, uuid.UUID](c.Context(), command)
	if err != nil {
		log.Printf("Command Failure: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Command failed",
			"details": err.Error(),
		})
	}

	//return data
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userId": result,
	})
}
