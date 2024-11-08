package Repository

import (
	"context"
	"fmt"

	"github.com/taskManagement/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	dbInstance mongo.Database
}

func TaskRepo(db mongo.Database) *Task {
	return &Task{
		dbInstance: db,
	}
}

func (s *Task) CreateTask(ctx context.Context, task Model.Task) (primitive.ObjectID, error) {
	coll := s.dbInstance.Collection("tasks")
	result, err := coll.InsertOne(ctx, task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("error converting insertedID to objectId")
	}
	return id, nil
}

func (s *Task) FetchAllTasks(ctx context.Context) ([]Model.Task, error) {
	coll := s.dbInstance.Collection("tasks")
	var tasks []Model.Task
	query := bson.D{}
	// var user Model.User
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Task) FetchTask(ctx context.Context, id primitive.ObjectID) (Model.Task, error) {
	coll := s.dbInstance.Collection("tasks")
	fmt.Print("jell")
	var task Model.Task
	query := bson.D{{Key: "_id", Value: id}}
	err := coll.FindOne(ctx, query).Decode(&task)
	if err != nil {
		return Model.Task{}, nil
	}
	return task, nil
}

func (s *Task) UpdateTask(ctx context.Context, id primitive.ObjectID, updateJson bson.M) error {
	coll := s.dbInstance.Collection("tasks")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{
		"$set": updateJson,
	}

	result, err := coll.UpdateOne(ctx, filter, update)

	if err != nil {
		return fmt.Errorf("error updating task")
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no task found with id")
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("task is up-to-date")
	}

	return nil

}

func (s *Task) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	coll := s.dbInstance.Collection("tasks")
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		return fmt.Errorf("error deleting task ")
	}

	return nil

}

func (s *Task) GetTaskTitle(ctx context.Context) (map[string]string, error) {
	coll := s.dbInstance.Collection("tasks")
	query := bson.D{}
	result := make(map[string]string)
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		result[task.Title] = task.Id.Hex()
	}
	return result, nil
}

func (s *Task) IsTaskExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	coll := s.dbInstance.Collection("tasks")
	query := bson.D{}
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return false, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Model.Task
		if err := cursor.Decode(&task); err != nil {
			return false, err
		}
		if task.Id == id {
			return true, nil
		}

	}
	return false, nil
}
