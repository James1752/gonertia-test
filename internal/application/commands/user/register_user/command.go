package commmands_user

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserCommand struct {
	UserID       uuid.UUID `validate:"required"`
	FirstName    string    `validate:"required"`
	LastName     string    `validate:"required"`
	Email        string    `validate:"required,email"`
	CreatedAtUtc time.Time `validate:"required,datetime"`
}

func NewRegisterUserCommand(firstName string, lastName string, email string) *RegisterUserCommand {
	return &RegisterUserCommand{
		UserID:       uuid.New(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAtUtc: time.Now().UTC(),
	}
}
