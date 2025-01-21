package fiber_application

import (
	"log"

	"github.com/James1752/gonertia-test/pkg/application"
	"github.com/James1752/gonertia-test/pkg/module"
	"github.com/gofiber/fiber/v2"
)

type FiberApplication struct {
	fiberInstance *fiber.App
	modules       []module.ModuleService
}

func NewFiberApplication(modules []module.ModuleService) application.Application {
	app := &FiberApplication{
		fiberInstance: fiber.New(),
		modules:       modules,
	}

	app.registerRoutes(app.fiberInstance)

	return app
}

func (fa *FiberApplication) registerRoutes(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pooo\n")
	})
}

// func forwardFiberRequestToModule(module module.Module) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		baseURL := module.GetHttpServerUrlBase()
// 		targetURL := baseURL + strings.TrimPrefix(c.Path(), "/api/v1")
// 		return proxy.Do(c, targetURL, &fasthttp.Client{})
// 	}
// }

func (a *FiberApplication) Run() error {
	return a.fiberInstance.Listen(":3000")
}

func (a *FiberApplication) RunModuleServices() {
	for _, module := range a.modules {
		go func() {
			if err := module.Start(); err != nil {
				log.Fatal(err)
			}
		}()
	}
}
