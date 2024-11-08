package Model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStore interface {
	CreateTask(ctx context.Context, task Task) (primitive.ObjectID, error)
	FetchAllTasks(ctx context.Context) ([]Task, error)
	FetchTask(ctx context.Context, id primitive.ObjectID) (Task, error)
	UpdateTask(ctx context.Context, id primitive.ObjectID, updateJson bson.M) error
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
	GetTaskTitle(ctx context.Context) (map[string]string, error)
	IsTaskExists(ctx context.Context, id primitive.ObjectID) (bool, error)
}

type Task struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	AssignedTo  string             `bson:"assignedTo" json:"assignedTo"`
	CreatedAt   string             `bson:"createdAt" json:"createdAt"`
}
