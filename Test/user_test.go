package Test

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taskManagement/Model"
	"github.com/taskManagement/Service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserRouter(t *testing.T) {
	taskID := primitive.NewObjectID()
	mockStore := &mockUserStore{
		Users: []Model.User{
			{
				Id:           taskID,
				Name:         "Test User",
				Email:        "testuser@example.com",
				UserName:     "testuser",
				Dob:          time.Hour * 24 * 365 * 30,
				PasswordHash: "hashedpassword",
				Gender:       "Male",
			},
			{
				Id:           taskID,
				Name:         "Test User",
				Email:        "testuser@example.com",
				UserName:     "testuser",
				Dob:          time.Hour * 24 * 365 * 30,
				PasswordHash: "hashedpassword",
				Gender:       "Male",
			},
		},
	}

	serviceHandler := Service.NewServiceHandler(mockStore, nil)
	app := fiber.New()

	t.Run("Sign Up", func(t *testing.T) {
		app.Post("/signUp", serviceHandler.SignUp)
		taskPayload := `{
			"fullName" : "Prem Kumar",
			"userName" : "prem310501@gmail.com",
			"password" : "cricket@guitar123"
		}`
		req := httptest.NewRequest("POST", "/signUp", bytes.NewReader([]byte(taskPayload)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Sign In", func(t *testing.T) {
		app.Post("/signIn", serviceHandler.SignIn)
		taskPayload := `{
			"username" : "prem310501@gmail.com",
			"password" : "cricket@guitar123"
		}`
		req := httptest.NewRequest("POST", "/signIn", bytes.NewReader([]byte(taskPayload)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("SignOut", func(t *testing.T) {
		app.Post("/signOut", serviceHandler.SignOut)
		req := httptest.NewRequest("POST", "/signOut", nil)
		req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDI0LTExLTA3VDE0OjU0OjUwLjk4NTM0MzY1NCswNTozMCJ9.zes7vMaoBTNLqHCgaDvDQRha4QyrJWPp5qOZig20H0E")
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	})

}

type mockUserStore struct {
	Users []Model.User
}

func (m *mockUserStore) GetUserDetails(ctx context.Context, id int) (Model.User, error) {
	return Model.User{}, nil
}

func (m *mockUserStore) CreateUser(ctx context.Context, user Model.User) error {
	m.Users = append(m.Users, user)
	return nil
}

func (m *mockUserStore) FetchUser(ctx context.Context, email string) Model.User {
	val := Model.User{PasswordHash: "$2a$10$XCmuzsPbaLtFQ2dDplOAsuT43jehc9RYa79YpOqZbQLmmH2VSfFWePASS"}
	val.Email = email
	return val
}

func (m *mockUserStore) GetUserName(ctx context.Context) (map[string]string, error) {
	return map[string]string{}, nil
}
