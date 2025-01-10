package domain_user

import "github.com/google/uuid"

type UserRepository interface {
	GetUserById(is uuid.UUID) (*User, error)
	CreateUser(*User) error
}
