package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taskManagement/Middleware"
	"github.com/taskManagement/Service"
)

type Handler struct {
	service *Service.ServiceHandler
}

func NewHandler(ser *Service.ServiceHandler) *Handler {
	return &Handler{
		service: ser,
	}
}

func (s *Handler) ProtectedHandler(app *fiber.App) {
	app.Post("/tasks", Middleware.JwtMiddleware, s.CreateTask)
	app.Get("/tasks", Middleware.JwtMiddleware, s.GetAllTask)
	app.Get("/tasks/:id", Middleware.JwtMiddleware, s.FetchTask)
	app.Put("/tasks/:id", Middleware.JwtMiddleware, s.UpdateTask)
	app.Delete("/tasks/:id", Middleware.JwtMiddleware, s.DeleteTask)

}

func (s *Handler) UnProtectedHandler(app *fiber.App) {
	app.Post("/signUp", s.SignUp)
	app.Post("/signIn", s.LogIn)
	app.Post("/signOut", s.SignOut)
}

func (s *Handler) LogIn(r *fiber.Ctx) error {
	return s.service.SignIn(r)
}

func (s *Handler) SignUp(r *fiber.Ctx) error {
	return s.service.SignUp(r)
}
func (s *Handler) SignOut(r *fiber.Ctx) error {
	return s.service.SignOut(r)
}

func (s *Handler) CreateTask(r *fiber.Ctx) error {
	return s.service.CreateTask(r)
}

func (s *Handler) GetAllTask(r *fiber.Ctx) error {
	return s.service.GetAllTask(r)
}

func (s *Handler) FetchTask(r *fiber.Ctx) error {
	return s.service.GetTask(r)
}

func (s *Handler) UpdateTask(r *fiber.Ctx) error {
	return s.service.UpdateTask(r)
}

func (s *Handler) DeleteTask(r *fiber.Ctx) error {
	return s.service.DeleteTask(r)
}
