package api_server

import (
	"log"

	controllers_user "github.com/James1752/gonertia-test/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	server := &Server{
		app: fiber.New(),
	}

	server.setupRoutes()

	return server
}

func (s *Server) StartServer() {
	log.Fatal(s.app.Listen(":3000"))
}

func (s *Server) setupRoutes() {
	api := s.app.Group("/api/v1")

	//User Controller
	controllers_user.NewUserController().RegisterRoutes(api)
}
