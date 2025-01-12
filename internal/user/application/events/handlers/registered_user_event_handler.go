package user_events

import (
	"context"
	"fmt"

	user_events "github.com/James1752/gonertia-test/internal/user/application/events"
)

type UserRegisteredEventHandler struct {
}

func NewUserRegisteredEventHandler() *UserRegisteredEventHandler {
	return &UserRegisteredEventHandler{}
}

func (c *UserRegisteredEventHandler) Handle(ctx context.Context, event *user_events.UserRegisteredEvent) error {
	fmt.Printf("User has been Registered: %s %s, %s\n", event.FirstName, event.LastName, event.Email)

	return nil
}
