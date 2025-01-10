package events_user

import (
	"context"
	"fmt"
)

type UserRegisteredEventHandler struct {
}

func NewUserRegisteredEventHandler() *UserRegisteredEventHandler {
	return &UserRegisteredEventHandler{}
}

func (c *UserRegisteredEventHandler) Handle(ctx context.Context, event *RegisterUserCommand) error {
	fmt.Printf("User has been Registered: %s %s, %s\n", event.FirstName, event.LastName, event.Email)

	return nil
}
