package server

import (
	"log"

	user_controller "github.com/James1752/gonertia-test/internal/user/presentation/api_controllers"
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
	user_controller.NewUserController().RegisterRoutes(api)
}
