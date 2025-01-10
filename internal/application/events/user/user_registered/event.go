package events_user

type RegisterUserCommand struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

func NewRegisterUserCommand(firstName string, lastName string, email string) *RegisterUserCommand {
	return &RegisterUserCommand{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
