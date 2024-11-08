package Repository

import (
	"context"

	"github.com/taskManagement/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	dbInstance mongo.Database
}

func NewUserRepo(db mongo.Database) *UserRepo {
	return &UserRepo{
		dbInstance: db,
	}
}

func (s *UserRepo) GetUserDetails(ctx context.Context, id int) (Model.User, error) {
	collection := s.dbInstance.Collection("user")
	filter := bson.D{{Key: "id", Value: id}}
	var result Model.User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *UserRepo) CreateUser(ctx context.Context, user Model.User) error {
	coll := s.dbInstance.Collection("user")
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepo) FetchUser(ctx context.Context, email string) Model.User {
	coll := s.dbInstance.Collection("user")
	query := bson.D{{Key: "email", Value: email}}
	var user Model.User
	coll.FindOne(ctx, query).Decode(&user)
	return user
}

func (s *UserRepo) GetUserName(ctx context.Context) (map[string]string, error) {
	coll := s.dbInstance.Collection("user")
	query := bson.D{}
	result := make(map[string]string)
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Model.User
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		result[task.UserName] = task.Name
	}
	return result, nil
}
