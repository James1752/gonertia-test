package main

import (
	"log"

	api_server "github.com/James1752/gonertia-test/internal/api"
	commmands_user "github.com/James1752/gonertia-test/internal/application/commands/user/register_user"
	events_user "github.com/James1752/gonertia-test/internal/application/events/user/user_registered"
	infrastructure_user "github.com/James1752/gonertia-test/internal/infrastructure/user"
	"github.com/mehdihadeli/go-mediatr"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	//Repositories

	userRepository := infrastructure_user.NewUserInMemoryRepository()

	//Handlers
	//Commands
	err := mediatr.RegisterRequestHandler(
		commmands_user.NewRegisterUserCommandHandler(userRepository),
	)
	if err != nil {
		log.Fatal(err)
	}

	//Queries

	//Events
	mediatr.RegisterNotificationHandler(
		events_user.NewUserRegisteredEventHandler(),
	)
	if err != nil {
		log.Fatal(err)
	}

	server := api_server.NewServer()
	server.StartServer()
}
