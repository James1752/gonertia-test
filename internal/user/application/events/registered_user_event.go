package user_events

type UserRegisteredEvent struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

func NewRegisterUserCommand(firstName string, lastName string, email string) *UserRegisteredEvent {
	return &UserRegisteredEvent{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
