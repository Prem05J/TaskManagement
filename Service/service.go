package Service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/taskManagement/Model"
	"github.com/taskManagement/Request"
	"github.com/taskManagement/Response"
	"github.com/taskManagement/Utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceHandler struct {
	userStore Model.UserStore
	taskStore Model.TaskStore
}

func NewServiceHandler(userStore Model.UserStore, taskStore Model.TaskStore) *ServiceHandler {
	return &ServiceHandler{
		userStore: userStore,
		taskStore: taskStore,
	}
}

func (s *ServiceHandler) SignUp(c *fiber.Ctx) error {
	var userReq Request.SignUpRequest
	var user Model.User

	if err := c.BodyParser(&userReq); err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, "Invalid Input data")
	}

	validator := validator.New()
	if err := validator.Struct(userReq); err != nil {
		return Utils.WriteFiberMap(c, fiber.StatusBadRequest, "error", err.Error())
	}

	users, err := s.userStore.GetUserName(c.Context())
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	if _, exists := users[userReq.Username]; exists {
		return Utils.WriteErrorJson(c, fiber.StatusBadGateway, "User already exists")
	}

	hashPassword, err := Utils.HashPassword(userReq.Password)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, "Error Hashing password")
	}

	user.PasswordHash = hashPassword
	user.Name = userReq.Fullname
	user.Email = userReq.Username
	user.UserName = userReq.Username

	s.userStore.CreateUser(context.Background(), user)

	token, err := Utils.GenerateJWT(user)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteFiberMap(c, fiber.StatusOK, "token", token)
}

func (s *ServiceHandler) SignIn(c *fiber.Ctx) error {
	var userReq Request.SignInRequest

	if err := c.BodyParser(&userReq); err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, "Invalid Request")
	}

	validator := validator.New()
	if err := validator.Struct(userReq); err != nil {
		return Utils.WriteFiberMap(c, fiber.StatusBadRequest, "error", err.Error())
	}

	user := s.userStore.FetchUser(c.Context(), userReq.Username)
	if user.Email == "" && user.PasswordHash == "" {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, "User Not Found")
	}

	if !Utils.VerifyPassword(user.PasswordHash, userReq.Password) {
		return Utils.WriteErrorJson(c, fiber.StatusUnauthorized, "Invalid Password")
	}

	token, err := Utils.GenerateJWT(user)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteFiberMap(c, fiber.StatusOK, "token", token)
}

func (s *ServiceHandler) SignOut(c *fiber.Ctx) error {
	var jwtSecretKey = []byte(Utils.GetEnv("SECRET_KEY", "your-secret-key"))
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return Utils.WriteErrorJson(c, fiber.StatusUnauthorized, "Authorication header missing")
	}
	tokenString := authHeader[len("Bearer "):]

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid Signing Method")
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		Utils.WriteErrorJson(c, fiber.StatusBadRequest, "Invalid Signing Method")
	}

	expirationTime := time.Now().Add(-time.Hour)

	claims := jwt.MapClaims{
		"exp": expirationTime,
	}

	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := expiredToken.SignedString(jwtSecretKey)

	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteFiberMap(c, fiber.StatusOK, "token", token)
}

func (s *ServiceHandler) CreateTask(c *fiber.Ctx) error {
	var task Model.Task
	var resp Response.TaskResponse
	if err := c.BodyParser(&task); err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, err.Error())
	}

	validator := validator.New()
	if err := validator.Struct(task); err != nil {
		return Utils.WriteFiberMap(c, fiber.StatusBadRequest, "error", err.Error())
	}

	users, err := s.taskStore.GetTaskTitle(c.Context())
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	if value, exists := users[task.Title]; exists {
		val := fiber.Map{"error": "Task already exists", "id": value}
		return Utils.WriteMap(c, fiber.StatusBadGateway, val)
	}

	task.CreatedAt = time.Now().String()

	id, err := s.taskStore.CreateTask(c.Context(), task)
	resp.Id = id.Hex()

	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteJson(c, fiber.StatusOK, resp)
}

func (s *ServiceHandler) GetAllTask(c *fiber.Ctx) error {
	resp, err := s.taskStore.FetchAllTasks(c.Context())
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}
	if resp == nil {
		return Utils.WriteFiberMap(c, fiber.StatusOK, "result", "No task available")
	}
	return Utils.WriteJson(c, fiber.StatusOK, resp)
}

func (s *ServiceHandler) GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Print(id)
	objId, err := primitive.ObjectIDFromHex(id)
	flag, _ := s.taskStore.IsTaskExists(c.Context(), objId)
	if !flag {
		val := fiber.Map{"error": "Task not Exists", "id": objId.Hex()}
		return Utils.WriteMap(c, fiber.StatusBadGateway, val)
	}
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, err.Error())
	}
	resp, err := s.taskStore.FetchTask(c.Context(), objId)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}
	return Utils.WriteJson(c, fiber.StatusOK, resp)
}

func (s *ServiceHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task Request.UpdateTaskRequest
	var resp Response.TaskResponse
	resp.Id = id
	if err := c.BodyParser(&task); err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, err.Error())
	}
	objId, err := primitive.ObjectIDFromHex(id)
	flag, _ := s.taskStore.IsTaskExists(c.Context(), objId)
	if !flag {
		val := fiber.Map{"error": "Task not Exists", "id": objId.Hex()}
		return Utils.WriteMap(c, fiber.StatusBadGateway, val)
	}

	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, err.Error())
	}

	fieldtoUpdate := bson.M{}

	if task.Status != "" {
		fieldtoUpdate["status"] = task.Status
	}
	if task.Description != "" {
		fieldtoUpdate["description"] = task.Description
	}
	if task.Title != "" {
		fieldtoUpdate["title"] = task.Title
	}

	if err := s.taskStore.UpdateTask(c.Context(), objId, fieldtoUpdate); err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteJson(c, fiber.StatusOK, resp)
}

func (s *ServiceHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusBadRequest, err.Error())
	}
	flag, _ := s.taskStore.IsTaskExists(c.Context(), objId)
	if !flag {
		val := fiber.Map{"error": "Task not Exists", "id": objId.Hex()}
		return Utils.WriteMap(c, fiber.StatusBadGateway, val)
	}
	err = s.taskStore.DeleteTask(c.Context(), objId)
	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusInternalServerError, err.Error())
	}

	return Utils.WriteJson(c, fiber.StatusOK, fiber.Map{"result": "Task Deleted"})

}
