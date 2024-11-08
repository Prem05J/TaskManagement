package application

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taskManagement/Handler"
	"github.com/taskManagement/Repository"
	"github.com/taskManagement/Service"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
	db   *mongo.Database
}

func NewAPIServer(addr string, db *mongo.Database) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()

	taskRepo := Repository.TaskRepo(*s.db)
	userRepo := Repository.NewUserRepo(*s.db)
	service := Service.NewServiceHandler(userRepo, taskRepo)
	handler := Handler.NewHandler(service)
	handler.ProtectedHandler(app)
	handler.UnProtectedHandler(app)

	log.Println("ListeningON", s.addr)
	return app.Listen(s.addr)

}
