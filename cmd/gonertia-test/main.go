package main

import (
	"log"

	"go.uber.org/dig"
)

// Define interfaces for the services/modules
type Service interface {
	Name() string
}

// Concrete implementation of ServiceA
type ServiceA struct{}

func (s *ServiceA) Name() string {
	return "Service A"
}

// Concrete implementation of ServiceB
type ServiceB struct{}

func (s *ServiceB) Name() string {
	return "Service B"
}

// Module registration for ServiceA
func NewServiceA() Service {
	return &ServiceA{}
}

// Module registration for ServiceB
func NewServiceB() Service {
	return &ServiceB{}
}

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	// //Create Modules
	// userModule := user.NewUserModule("localhost", "3001")
	// if err := userModule.StartHttpServer(); err != nil {
	// 	log.Fatal(err)
	// }

	// //Create the root api router
	// app := fiber.New()
	// apiRouter := app.Group("/api/v1")

	// //Forward all /api/v1/user requests to the user module server
	// apiRouter.All("/user/*", forwardFiberRequestToModule(userModule))
	// //Forward all /api/v1/test requests to the user module server
	// apiRouter.All("/test/*", forwardFiberRequestToModule(userModule))

	// //Start Server
	// log.Fatal(app.Listen(":3000"))

	// app := application.NewBoostrapApplication(fiber_application.NewFiberApplication)
	// app.RegisterModuleServices(func(container dig.Container) module.ModuleService {
	// 	return user.NewUserModule(container, &user.UserModuleConfig{
	// 		Host: "localhost",
	// 		Port: "3001",
	// 	})
	// })

	// if err := app.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	container := dig.New()

	// Register ServiceA and ServiceB as part of a "services" group
	type ServiceGroupOut struct {
		dig.Out

		Service Service `group:"services"`
	}

	if err := container.Provide(func() ServiceGroupOut {
		return ServiceGroupOut{
			Service: NewServiceA(),
		}
	}); err != nil {
		log.Fatal("Error:", err)
	}
	if err := container.Provide(func() ServiceGroupOut {
		return ServiceGroupOut{
			Service: NewServiceB(),
		}
	}); err != nil {
		log.Fatal("Error:", err)
	}

	// Invoke the container to get all services
	type ServiceGroupIn struct {
		dig.In
		Service []Service `group:"services"`
	}
	err := container.Invoke(func(services ServiceGroupIn) {
		for _, service := range services.Service {
			log.Println(service.Name())
		}
	})

	if err != nil {
		log.Fatal("Error:", err)
	}

}
