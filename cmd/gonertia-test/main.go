package main

import (
	"log"

	server "github.com/James1752/gonertia-test/internal"
	user_commands "github.com/James1752/gonertia-test/internal/user/application/commands/handlers"
	user_events "github.com/James1752/gonertia-test/internal/user/application/events/handlers"
	user_domain "github.com/James1752/gonertia-test/internal/user/domain"
	user_infrastructure "github.com/James1752/gonertia-test/internal/user/infrastructure"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

// SetupContainer sets up the DI container and provides dependencies.
func SetupContainer() *dig.Container {
	container := dig.New()

	// Provide dependencies
	container.Provide(func() user_domain.UserRepository {
		return user_infrastructure.NewUserInMemoryRepository()
	})
	container.Provide(user_commands.NewRegisterUserCommandHandler)
	container.Provide(user_events.NewUserRegisteredEventHandler)
	container.Provide(server.NewServer)

	return container
}

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	container := SetupContainer()
	err := container.Invoke(func(
		//Repositories
		userRepository user_domain.UserRepository,
		//Commands
		registerUserHandler *user_commands.RegisterUserCommandHandler,
		//Events
		userRegisteredEvent *user_events.UserRegisteredEventHandler,
		//Server
		server *server.Server,
	) {

		//Commands
		if err := mediatr.RegisterRequestHandler(registerUserHandler); err != nil {
			log.Fatal(err)
		}

		//Events
		if err := mediatr.RegisterNotificationHandler(userRegisteredEvent); err != nil {
			log.Fatal(err)
		}

		server.StartServer()
	})

	if err != nil {
		log.Fatal(err)
	}
}
