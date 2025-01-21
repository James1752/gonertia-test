package user

import (
	"fmt"
	"log"

	user_commands "github.com/James1752/gonertia-test/internal/user/application/commands/handlers"
	user_events "github.com/James1752/gonertia-test/internal/user/application/events/handlers"
	user_domain "github.com/James1752/gonertia-test/internal/user/domain"
	user_infrastructure "github.com/James1752/gonertia-test/internal/user/infrastructure"
	user_api_handlers "github.com/James1752/gonertia-test/internal/user/presentation/api_handlers"
	"github.com/James1752/gonertia-test/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type UserModuleConfig struct {
	Host string
	Port string
}
type UserModule struct {
	name      string
	config    *UserModuleConfig
	container dig.Container
}

func NewUserModule(container dig.Container, config *UserModuleConfig) *UserModule {
	module := &UserModule{
		name:      "user-module",
		config:    config,
		container: container,
	}

	if err := module.registerServices(); err != nil {
		log.Fatal(err)
	}

	if err := module.registerHandlers(); err != nil {
		log.Fatal(err)
	}

	return module
}

func (m *UserModule) Start() error {
	log.Printf("MADE IT HERE \n\n")
	return nil
}

func (m *UserModule) GetHttpServerUrlBase() string {
	return fmt.Sprintf("http://%s:%s/%s", m.config.Host, m.config.Port, m.name)
}

func (m *UserModule) StartHttpServer() error {
	return m.container.Invoke(func(
		controllers []api.FiberHandler,
	) {
		app := fiber.New()
		moduleRouter := app.Group("/" + m.name)

		for _, controller := range controllers {
			controller.RegisterRoutes(moduleRouter)
		}

		go func() {
			log.Printf("Starting %s module server on %s", m.name, m.GetHttpServerUrlBase())
			log.Fatal(app.Listen(m.config.Host + ":" + m.config.Port))
		}()
	})
}

func (m *UserModule) registerServices() error {
	providers := []func() error{
		//Handlers
		func() error {
			return m.container.Provide(func() []api.FiberHandler {
				return []api.FiberHandler{
					user_api_handlers.NewUserApiHandler(),
					user_api_handlers.NewTestApiHandler(),
				}
			})
		},

		//Repositories
		func() error {
			return m.container.Provide(func() user_domain.UserRepository {
				return user_infrastructure.NewUserInMemoryRepository()
			})
		},
	}

	for _, provide := range providers {
		if err := provide(); err != nil {
			return err
		}
	}

	return nil
}

func (m *UserModule) registerHandlers() error {
	return m.container.Invoke(func(
		userRepository user_domain.UserRepository,
	) {
		handlers := []func() error{
			//Queries

			////////////////////////////////////////////////////////////////////////////////////////////////

			//Commands
			//Register New User
			func() error {
				command := user_commands.NewRegisterUserCommandHandler(userRepository)
				return registerRequestHandler(command)
			},

			////////////////////////////////////////////////////////////////////////////////////////////////

			//Events
			//User Registered
			func() error {
				event := user_events.NewUserRegisteredEventHandler()
				return registerNotificationHandler(event)
			},
		}

		for _, handler := range handlers {
			handler()
		}
	})
}

func registerRequestHandler[TRequest any, TResponse any](handler mediatr.RequestHandler[TRequest, TResponse]) error {
	if err := mediatr.RegisterRequestHandler(handler); err != nil {
		return err
	}
	return nil
}

func registerNotificationHandler[TEvent any](handler mediatr.NotificationHandler[TEvent]) error {
	if err := mediatr.RegisterNotificationHandler(handler); err != nil {
		return err
	}
	return nil
}
