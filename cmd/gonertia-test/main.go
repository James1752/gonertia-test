package main

import (
	"log"
	"strings"

	"github.com/James1752/gonertia-test/internal/user"
	"github.com/James1752/gonertia-test/pkg/module"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func forwardFiberRequestToModule(module module.Module) fiber.Handler {
	return func(c *fiber.Ctx) error {
		baseURL := module.GetHttpServerUrlBase()
		targetURL := baseURL + strings.TrimPrefix(c.Path(), "/api/v1")
		return proxy.Do(c, targetURL, &fasthttp.Client{})
	}
}

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	//Create Modules
	userModule := user.NewUserModule("localhost", "3001")
	if err := userModule.StartHttpServer(); err != nil {
		log.Fatal(err)
	}

	//Create the root api router
	app := fiber.New()
	apiRouter := app.Group("/api/v1")

	//Forward all /api/v1/user requests to the user module server
	apiRouter.All("/user/*", forwardFiberRequestToModule(userModule))
	//Forward all /api/v1/test requests to the user module server
	apiRouter.All("/test/*", forwardFiberRequestToModule(userModule))

	//Start Server
	log.Fatal(app.Listen(":3000"))
}
