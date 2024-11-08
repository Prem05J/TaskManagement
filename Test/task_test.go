package Test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taskManagement/Model"
	"github.com/taskManagement/Service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var mockTasks = GetMockTasks()

func TestTaskRouter(t *testing.T) {
	taskID := primitive.NewObjectID()
	mockStore := &mockTaskStore{
		Tasks: []Model.Task{
			{
				Id:          taskID,
				Title:       "Complete the proposal",
				Description: "Finalize the project proposal for the new feature",
				Status:      "In Progress",
				AssignedTo:  "Alice",
				CreatedAt:   "2024-10-15T10:30:00Z",
			},
			{
				Id:          taskID,
				Title:       "Complete the proposal",
				Description: "Finalize the project proposal for the new feature",
				Status:      "In Progress",
				AssignedTo:  "Alice",
				CreatedAt:   "2024-10-15T10:30:00Z",
			},
		},
	}

	serviceHandler := Service.NewServiceHandler(nil, mockStore)
	app := fiber.New()

	t.Run("Create Task", func(t *testing.T) {
		app.Post("/tasks", serviceHandler.CreateTask)
		taskPayload := `{
			"title" : "UI Desing",
			"status": "Pending",
			"assignedTo": "Charlie"
		}`
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte(taskPayload)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var responseMap map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)
		_, exists := responseMap["id"]
		assert.True(t, exists, "Response does not contain task ID")
	})

	t.Run("Get All Task", func(t *testing.T) {
		app.Get("/tasks", serviceHandler.GetAllTask)

		req := httptest.NewRequest("GET", "/tasks", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var responseMap []Model.Task
		err = json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)
	})

	t.Run("Get Task", func(t *testing.T) {
		app.Get("/tasks/:id", serviceHandler.GetTask)

		req := httptest.NewRequest("GET", fmt.Sprintf("/tasks/%s", taskID.Hex()), nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var responseMap map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)

	})

	t.Run("Delete Task", func(t *testing.T) {
		app.Delete("/tasks/:id", serviceHandler.DeleteTask)

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/tasks/%s", taskID.Hex()), nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var responseMap map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)
	})

	t.Run("Update Task", func(t *testing.T) {
		app.Put("/tasks/:id", serviceHandler.UpdateTask)
		taskPayload := `{
			"title": "Sidhu"
		}`
		req := httptest.NewRequest("PUT", fmt.Sprintf("/tasks/%s", taskID.Hex()), bytes.NewReader([]byte(taskPayload)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var responseMap map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)
	})

}

type mockTaskStore struct {
	Tasks []Model.Task
}

func (m *mockTaskStore) CreateTask(ctx context.Context, task Model.Task) (primitive.ObjectID, error) {
	task.Id = primitive.NewObjectID()
	m.Tasks = append(m.Tasks, task)
	return task.Id, nil
}

func (m *mockTaskStore) FetchAllTasks(ctx context.Context) ([]Model.Task, error) {
	return m.Tasks, nil
}

func (m *mockTaskStore) FetchTask(ctx context.Context, id primitive.ObjectID) (Model.Task, error) {
	for _, task := range m.Tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return Model.Task{}, fmt.Errorf("task not found")
}

func (m *mockTaskStore) UpdateTask(ctx context.Context, id primitive.ObjectID, updateJson bson.M) error {
	for i, task := range m.Tasks {
		if task.Id == id {
			if title, ok := updateJson["title"]; ok {
				m.Tasks[i].Title = title.(string)
			}
			if description, ok := updateJson["description"]; ok {
				m.Tasks[i].Description = description.(string)
			}
			if status, ok := updateJson["status"]; ok {
				m.Tasks[i].Status = status.(string)
			}
			if assignedTo, ok := updateJson["assignedTo"]; ok {
				m.Tasks[i].AssignedTo = assignedTo.(string)
			}
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (m *mockTaskStore) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	for i, task := range m.Tasks {
		if task.Id == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (m *mockTaskStore) GetTaskTitle(ctx context.Context) (map[string]string, error) {
	return map[string]string{}, nil
}

func (m *mockTaskStore) IsTaskExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	return true, nil
}
