package user_commands

import (
	"context"

	user_commands "github.com/James1752/gonertia-test/internal/user/application/commands"
	user_events "github.com/James1752/gonertia-test/internal/user/application/events"
	user_domain "github.com/James1752/gonertia-test/internal/user/domain"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type RegisterUserCommandHandler struct {
	userRepository user_domain.UserRepository
}

func NewRegisterUserCommandHandler(userRepository user_domain.UserRepository) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{userRepository: userRepository}
}

func (c *RegisterUserCommandHandler) Handle(ctx context.Context, command *user_commands.RegisterUserCommand) (uuid.UUID, error) {
	//Create the user entity
	user := &user_domain.User{
		UserID:       command.UserID,
		FirstName:    command.FirstName,
		LastName:     command.LastName,
		Email:        command.Email,
		CreatedAtUtc: command.CreatedAtUtc,
		UpdatedAtUtc: command.CreatedAtUtc,
	}

	//Create the user
	err := c.userRepository.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}

	//Publish event
	userRegisteredEvent := user_events.NewRegisterUserCommand(user.FirstName, user.LastName, user.Email)
	err = mediatr.Publish(ctx, userRegisteredEvent)
	if err != nil {
		return uuid.Nil, err
	}

	return user.UserID, nil
}
