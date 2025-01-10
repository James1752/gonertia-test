package commmands_user

import (
	"context"

	events_user "github.com/James1752/gonertia-test/internal/application/events/user/user_registered"
	user_domain "github.com/James1752/gonertia-test/internal/domain/user"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type RegisterUserCommandHandler struct {
	userRepository user_domain.UserRepository
}

func NewRegisterUserCommandHandler(userRepository user_domain.UserRepository) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{userRepository: userRepository}
}

func (c *RegisterUserCommandHandler) Handle(ctx context.Context, command *RegisterUserCommand) (uuid.UUID, error) {
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
	userRegisteredEvent := events_user.NewRegisterUserCommand(user.FirstName, user.LastName, user.Email)
	err = mediatr.Publish(ctx, userRegisteredEvent)
	if err != nil {
		return uuid.Nil, err
	}

	return user.UserID, nil
}
