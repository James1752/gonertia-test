package domain_user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `validate:"required"`
	FirstName    string    `validate:"required"`
	LastName     string    `validate:"required"`
	Email        string    `validate:"required,email"`
	CreatedAtUtc time.Time `validate:"required,datetime"`
	UpdatedAtUtc time.Time `validate:"required,datetime"`
}
