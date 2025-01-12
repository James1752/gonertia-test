package user

import (
	"fmt"
	"log"

	user_commands "github.com/James1752/gonertia-test/internal/user/application/commands/handlers"
	user_events "github.com/James1752/gonertia-test/internal/user/application/events/handlers"
	user_domain "github.com/James1752/gonertia-test/internal/user/domain"
	user_infrastructure "github.com/James1752/gonertia-test/internal/user/infrastructure"
	user_controllers "github.com/James1752/gonertia-test/internal/user/presentation/api_controllers"
	"github.com/James1752/gonertia-test/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/dig"
)

type UserModule struct {
	name      string
	host      string
	port      string
	container *dig.Container
}

func NewUserModule(host string, port string) *UserModule {
	module := &UserModule{
		name:      "user-module",
		host:      host,
		port:      port,
		container: dig.New(),
	}

	if err := module.registerServices(); err != nil {
		log.Fatal(err)
	}

	if err := module.registerHandlers(); err != nil {
		log.Fatal(err)
	}

	return module
}

func (m *UserModule) GetHttpServerUrlBase() string {
	return fmt.Sprintf("http://%s:%s/%s", m.host, m.port, m.name)
}

func (m *UserModule) StartHttpServer() error {
	return m.container.Invoke(func(
		controllers []api.FiberController,
		apiControllers []api.FiberApiController,
	) {
		app := fiber.New()
		moduleRouter := app.Group("/" + m.name)

		for _, controller := range controllers {
			controller.RegisterRoutes(moduleRouter)
		}
		for _, controller := range apiControllers {
			controller.RegisterRoutes(moduleRouter)
		}

		go func() {
			log.Printf("Starting %s module server on %s", m.name, m.GetHttpServerUrlBase())
			log.Fatal(app.Listen(m.host + ":" + m.port))
		}()
	})
}

func (m *UserModule) registerServices() error {
	providers := []func() error{
		//Controllers
		func() error {
			return m.container.Provide(func() []api.FiberController {
				return []api.FiberController{}
			})
		},
		//API Controllers
		func() error {
			return m.container.Provide(func() []api.FiberApiController {
				return []api.FiberApiController{
					user_controllers.NewUserApiController(),
					user_controllers.NewTestApiController(),
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
