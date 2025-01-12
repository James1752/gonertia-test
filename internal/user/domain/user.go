package user_domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	CreatedAtUtc time.Time
	UpdatedAtUtc time.Time
}
